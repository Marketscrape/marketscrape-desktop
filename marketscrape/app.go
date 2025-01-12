package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/iunary/fakeuseragent"
)

// App struct
type App struct {
	ctx context.Context
}

// Marketplace listing struct
type MarketplaceListingResult struct {
	Data  map[string]interface{} `json:"data"`
	Error string                 `json:"error,omitempty"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func GetFieldAsMap(field string, jsonObj map[string]interface{}) (map[string]interface{}, error) {
	jsonField, ok := jsonObj[field].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("field '%s' not found or invalid", field)
	}
	return jsonField, nil
}

func ParseMarketplaceListing(rawJson map[string]interface{}) (map[string]interface{}, error) {
	data, err := GetFieldAsMap("data", rawJson)
	if err != nil {
		return nil, fmt.Errorf("error accessing 'data': %w", err)
	}

	viewer, err := GetFieldAsMap("viewer", data)
	if err != nil {
		return nil, fmt.Errorf("error accessing 'viewer': %w", err)
	}

	marketplaceProductDetailsPage, err := GetFieldAsMap("marketplace_product_details_page", viewer)
	if err != nil {
		return nil, fmt.Errorf("error accessing 'marketplace_product_details_page': %w", err)
	}

	return marketplaceProductDetailsPage, nil
}

func (a *App) GetMarketplaceListing(id string) MarketplaceListingResult {
	r, _ := regexp.Compile("^[0-9]{15,16}$")
	if !r.MatchString(id) {
		return MarketplaceListingResult{
			Error: "Error parsing Marketplace listing ID",
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Use url.Values for proper payload construction
	payload := url.Values{}
	payload.Set("variables", fmt.Sprintf(`{"targetId":"%s"}`, id))
	payload.Set("doc_id", "7616889011758848")

	req, err := http.NewRequest("POST", "https://www.facebook.com/api/graphql/", strings.NewReader(payload.Encode()))
	if err != nil {
		return MarketplaceListingResult{
			Error: "Error creating request",
		}
	}

	req.Header.Set("User-Agent", fakeuseragent.RandomUserAgent())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-CA,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-FB-Friendly-Name", "MarketplacePDPContainerQuery")
	req.Header.Set("Origin", "https://www.facebook.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.facebook.com/marketplace/?ref=app_tab")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Priority", "u=1")
	req.Header.Set("TE", "trailers")

	res, err := client.Do(req)
	if err != nil {
		return MarketplaceListingResult{
			Error: "Error performing request",
		}
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return MarketplaceListingResult{
			Error: "Error reading response body",
		}
	}

	var rawJson map[string]interface{}
	err = json.Unmarshal(bodyBytes, &rawJson)
	if err != nil {
		return MarketplaceListingResult{
			Error: "Error unmarshalling response",
		}
	}

	parsedData, err := ParseMarketplaceListing(rawJson)
	if err != nil {
		return MarketplaceListingResult{
			Error: err.Error(),
		}
	}

	// Save the raw JSON to a file
	err = SaveJSONToFile("output.json", rawJson)
	if err != nil {
		fmt.Println("Error saving JSON to file:", err)
	}

	return MarketplaceListingResult{
		Data: parsedData,
	}
}

// SaveJSONToFile saves a JSON object to a file
func SaveJSONToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
