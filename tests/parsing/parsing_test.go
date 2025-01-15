package parsing

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

type Root struct {
	ProductDetailsType                 string                             `json:"product_details_type"`
	Target                             Target                             `json:"target"`
	MarketplaceListingRenderableTarget MarketplaceListingRenderableTarget `json:"marketplace_listing_renderable_target"`
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
}

type MarketplaceListingRenderableTarget struct {
	SEOCategories SEOVirtualCategory `json:"seo_virtual_category"`
}

type SEOVirtualCategory struct {
	ID           string `json:"id"`
	TaxonomyPath []struct {
		ID      string `json:"id"`
		SEOInfo struct {
			ID     string `json:"id"`
			SEOUrl string `json:"seo_url"`
		} `json:"seo_info"`
	} `json:"taxonomy_path"`
}

type ListingPhoto struct {
	ID                   string `json:"id"`
	AccessibilityCaption string `json:"accessibility_caption"`
	Image                struct {
		URI    string `json:"uri"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"image"`
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

func ParseMarketplaceListing(filename string) (*Root, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytesValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var root Root
	err = json.Unmarshal(bytesValue, &root)
	if err != nil {
		return nil, err
	}

	return &root, nil
}

func TestParseListing(t *testing.T) {
	root, err := ParseMarketplaceListing("product.json")
	if err != nil {
		t.Fatalf("Failed to parse vehicle listing: %v", err)
	}

	fmt.Print(root.Target.AttributeData)
}
