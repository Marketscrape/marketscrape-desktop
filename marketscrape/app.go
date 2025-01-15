package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/iunary/fakeuseragent"
)

// ---------------------------------------------------------
// SECTION: App
// ---------------------------------------------------------
type App struct {
	ctx context.Context
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

// ---------------------------------------------------------
// SECTION: Base struct, also applies to Products
// ---------------------------------------------------------
type Root struct {
	ProductDetailsType                 string                             `json:"product_details_type"`
	Target                             Target                             `json:"target"`
	MarketplaceListingRenderableTarget MarketplaceListingRenderableTarget `json:"marketplace_listing_renderable_target,omitempty"`
}

type Target struct {
	ID                  string             `json:"id"`
	Title               string             `json:"marketplace_listing_title"`
	ListingPhotos       []ListingPhoto     `json:"listing_photos"`
	PreRecordedVideos   []PreRecordedVideo `json:"pre_recorded_videos"`
	ListingPrice        ListingPrice       `json:"listing_price"`
	LocationText        LocationText       `json:"location_text"`
	RedactedDescription Description        `json:"redacted_description"`
	CreationTime        int64              `json:"creation_time"`
	AttributeData       []Attribute        `json:"attribute_data"`
	CommerceData        Commerce           `json:"commerce_badges_info"`
	VehicleData         *VehicleData       `json:"vehicle_data,omitempty"`
}

type ListingPhoto struct {
	ID                   string       `json:"id"`
	AccessibilityCaption string       `json:"accessibility_caption"`
	Image                ImageDetails `json:"image"`
}

type SEOVirtualCategory struct {
	ID           string             `json:"id"`
	TaxonomyPath []TaxonomyPathItem `json:"taxonomy_path"`
}

type PreRecordedVideo struct {
	ID          string `json:"id"`
	PlayableURL string `json:"playable_url"`
}

type ListingPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type LocationText struct {
	Text string `json:"text"`
}

type Description struct {
	Text string `json:"text"`
}

type Attribute struct {
	AttributeName string `json:"attribute_name"`
	Label         string `json:"label"`
	Value         string `json:"value"`
}

type Commerce struct {
	CommerceSummary string `json:"source_summary"`
}

type MarketplaceListingRenderableTarget struct {
	SEOCategories SEOVirtualCategory `json:"seo_virtual_category,omitempty"`
}

type ImageDetails struct {
	URI    string `json:"uri"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SEOInfo struct {
	ID     string `json:"id"`
	SEOUrl string `json:"seo_url"`
}

type TaxonomyPathItem struct {
	ID      string  `json:"id"`
	SEOInfo SEOInfo `json:"seo_info"`
}

// ---------------------------------------------------------
// SECTION: Vehicles struct
// ---------------------------------------------------------
type VehicleData struct {
	VehicleCarfaxReport                 any                   `json:"vehicle_carfax_report"`
	VehicleCondition                    any                   `json:"vehicle_condition"`
	VehicleExteriorColor                string                `json:"vehicle_exterior_color"`
	VehicleFeatures                     []any                 `json:"vehicle_features"`
	VehicleFuelType                     any                   `json:"vehicle_fuel_type"`
	VehicleIdentificationNumber         any                   `json:"vehicle_identification_number"`
	VehicleInteriorColor                string                `json:"vehicle_interior_color"`
	VehicleIsPaidOff                    any                   `json:"vehicle_is_paid_off"`
	VehicleMakeDisplayName              string                `json:"vehicle_make_display_name"`
	VehicleModelDisplayName             string                `json:"vehicle_model_display_name"`
	VehicleNumberOfOwners               any                   `json:"vehicle_number_of_owners"`
	VehicleOdometerData                 VehicleOdometerData   `json:"vehicle_odometer_data"`
	VehicleRegistrationPlateInformation any                   `json:"vehicle_registration_plate_information"`
	VehicleSellerType                   string                `json:"vehicle_seller_type"`
	VehicleSpecifications               VehicleSpecifications `json:"vehicle_specifications"`
	VehicleTitleStatus                  any                   `json:"vehicle_title_status"`
	VehicleTransmissionType             string                `json:"vehicle_transmission_type"`
	VehicleTrimDisplayName              any                   `json:"vehicle_trim_display_name"`
	VehicleWebsiteLink                  any                   `json:"vehicle_website_link"`
}

type VehicleOdometerData struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type VehicleSpecifications struct {
	Co2Emissions            any `json:"co2_emissions"`
	EngineSize              any `json:"engine_size"`
	GasMileageCity          any `json:"gas_mileage_city"`
	GasMileageCombined      any `json:"gas_mileage_combined"`
	GasMileageHighway       any `json:"gas_mileage_highway"`
	HorsePower              any `json:"horse_power"`
	SafetyRatingFront       any `json:"safety_rating_front"`
	SafetyRatingOverall     any `json:"safety_rating_overall"`
	SafetyRatingRollover    any `json:"safety_rating_rollover"`
	SafetyRatingSide        any `json:"safety_rating_side"`
	SafetyRatingSideBarrier any `json:"safety_rating_side_barrier"`
}

func GetFieldAsMap(field string, jsonObj map[string]interface{}) (map[string]interface{}, error) {
	jsonField, ok := jsonObj[field].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("field '%s' not found or invalid", field)
	}
	return jsonField, nil
}

func ParseMarketplaceListing(rawJson map[string]interface{}) (*Root, error) {
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

	pageBytes, err := json.Marshal(marketplaceProductDetailsPage)
	if err != nil {
		return nil, fmt.Errorf("error marshalling 'marketplace_product_details_page': %w", err)
	}

	var root Root
	err = json.Unmarshal(pageBytes, &root)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling to Root struct: %w", err)
	}

	// We're dealing with a vehicle
	if root.ProductDetailsType == "AUTOS_VEHICLE" {
		target, err := GetFieldAsMap("target", marketplaceProductDetailsPage)
		if err != nil {
			return nil, fmt.Errorf("error accessing 'viewer': %w", err)
		}

		targetBytes, err := json.Marshal(target)
		if err != nil {
			return nil, err
		}

		var vehicle VehicleData
		err = json.Unmarshal(targetBytes, &vehicle)
		if err != nil {
			return nil, err
		}

		root.Target.VehicleData = &vehicle
	}

	return &root, nil
}

type Response struct {
	Data  *Root   `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}

func (a *App) GetMarketplaceListing(id string) (*Root, error) {
	r, _ := regexp.Compile("^[0-9]{15,16}$")
	if !r.MatchString(id) {
		return nil, fmt.Errorf("invalid Marketplace listing ID: %s", id)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	payload := url.Values{}
	payload.Set("variables", fmt.Sprintf(`{"targetId":"%s"}`, id))
	payload.Set("doc_id", "7616889011758848")

	req, err := http.NewRequest("POST", "https://www.facebook.com/api/graphql/", strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
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
		return nil, fmt.Errorf("error performing request: %v", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var rawJson map[string]interface{}
	err = json.Unmarshal(bodyBytes, &rawJson)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	root, err := ParseMarketplaceListing(rawJson)
	if err != nil {
		return nil, fmt.Errorf("error parsing marketplace listing: %v", err)
	}

	return root, nil
}
