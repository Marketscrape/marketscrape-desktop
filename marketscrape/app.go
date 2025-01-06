package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iunary/fakeuseragent"
)

// App struct
type App struct {
	ctx context.Context
}

// Marketplace listing struct
type MarketplaceListing struct {
	Title        string   `json:"title"`
	Price        Price    `json:"price"`
	Description  string   `json:"description"`
	Location     Location `json:"location"`
	Condition    string   `json:"condition"`
	Brand        string   `json:"brand"`
	IsShipping   bool     `json:"is_shipping_offered"`
	Category     Category `json:"category"`
	CreationTime int64    `json:"creation_time"`
	IsSold       bool     `json:"is_sold"`
	Photos       []Photo  `json:"photos,omitempty"`
}

type Price struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	Formatted string `json:"formatted_amount"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Text      string  `json:"text"`
}

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Photo struct {
	URI    string `json:"uri"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
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

func ParseMarketplaceResponse(data map[string]interface{}) (*MarketplaceListing, error) {
	viewer, ok := data["data"].(map[string]interface{})["viewer"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response structure: missing viewer data")
	}

	details := viewer["marketplace_product_details_page"].(map[string]interface{})
	target := details["target"].(map[string]interface{})

	var condition, brand string
	if attrs, ok := target["attribute_data"].([]interface{}); ok {
		for _, attr := range attrs {
			attrMap := attr.(map[string]interface{})
			switch attrMap["attribute_name"] {
			case "Condition":
				condition = attrMap["label"].(string)
			case "Brand":
				brand = attrMap["label"].(string)
			}
		}
	}

	var photos []Photo
	if photosList, ok := target["listing_photos"].([]interface{}); ok {
		for _, p := range photosList {
			photoMap := p.(map[string]interface{})
			if img, ok := photoMap["image"].(map[string]interface{}); ok {
				photos = append(photos, Photo{
					URI:    img["uri"].(string),
					Width:  int(img["width"].(float64)),
					Height: int(img["height"].(float64)),
				})
			}
		}
	}

	listing := &MarketplaceListing{
		Title: target["marketplace_listing_title"].(string),
		Price: Price{
			Amount:    target["listing_price"].(map[string]interface{})["amount"].(string),
			Currency:  target["listing_price"].(map[string]interface{})["currency"].(string),
			Formatted: target["listing_price"].(map[string]interface{})["formatted_amount_zeros_stripped"].(string),
		},
		Description: target["redacted_description"].(map[string]interface{})["text"].(string),
		Location: Location{
			Latitude:  target["location"].(map[string]interface{})["latitude"].(float64),
			Longitude: target["location"].(map[string]interface{})["longitude"].(float64),
			Text:      target["location_text"].(map[string]interface{})["text"].(string),
		},
		Condition:    condition,
		Brand:        brand,
		IsShipping:   target["is_shipping_offered"].(bool),
		CreationTime: int64(target["creation_time"].(float64)),
		IsSold:       target["is_sold"].(bool),
		Category: Category{
			Name: target["marketplaceListingRenderableIfLoggedOut"].(map[string]interface{})["marketplace_listing_category_name"].(string),
			Slug: target["marketplaceListingRenderableIfLoggedOut"].(map[string]interface{})["marketplace_listing_category"].(map[string]interface{})["slug"].(string),
		},
		Photos: photos,
	}

	return listing, nil
}

func (a *App) GetMarketplaceListing(id string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	payload := fmt.Sprintf(`&variables={"targetId":"%s"}&doc_id=7616889011758848`, id)
	req, err := http.NewRequest("POST", "https://www.facebook.com/api/graphql/", strings.NewReader(url.PathEscape(payload)))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	headers := map[string]string{
		"User-Agent":         fakeuseragent.RandomUserAgent(),
		"Accept":             "*/*",
		"Accept-Language":    "en-CA,en-US;q=0.7,en;q=0.3",
		"Content-Type":       "application/x-www-form-urlencoded",
		"X-FB-Friendly-Name": "MarketplacePDPContainerQuery",
		"Origin":             "https://www.facebook.com",
		"Referer":            "https://www.facebook.com/marketplace/?ref=app_tab",
		"Connection":         "keep-alive",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"Priority":           "u=1",
		"TE":                 "trailers",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	listing, err := ParseMarketplaceResponse(result)
	if err != nil {
		return "", fmt.Errorf("error parsing marketplace data: %v", err)
	}

	jsonData, err := json.MarshalIndent(listing, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %v", err)
	}

	return string(jsonData), nil
}
