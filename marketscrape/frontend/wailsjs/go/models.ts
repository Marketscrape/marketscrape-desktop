export namespace main {
	
	export class Attribute {
	    attribute_name: string;
	    label: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Attribute(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.attribute_name = source["attribute_name"];
	        this.label = source["label"];
	        this.value = source["value"];
	    }
	}
	export class Commerce {
	    source_summary: string;
	
	    static createFrom(source: any = {}) {
	        return new Commerce(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source_summary = source["source_summary"];
	    }
	}
	export class Description {
	    text: string;
	
	    static createFrom(source: any = {}) {
	        return new Description(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.text = source["text"];
	    }
	}
	export class ImageDetails {
	    uri: string;
	    height: number;
	    width: number;
	
	    static createFrom(source: any = {}) {
	        return new ImageDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uri = source["uri"];
	        this.height = source["height"];
	        this.width = source["width"];
	    }
	}
	export class ListingPhoto {
	    id: string;
	    accessibility_caption: string;
	    image: ImageDetails;
	
	    static createFrom(source: any = {}) {
	        return new ListingPhoto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.accessibility_caption = source["accessibility_caption"];
	        this.image = this.convertValues(source["image"], ImageDetails);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ListingPrice {
	    amount: string;
	    currency: string;
	
	    static createFrom(source: any = {}) {
	        return new ListingPrice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.amount = source["amount"];
	        this.currency = source["currency"];
	    }
	}
	export class LocationText {
	    text: string;
	
	    static createFrom(source: any = {}) {
	        return new LocationText(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.text = source["text"];
	    }
	}
	export class SEOInfo {
	    id: string;
	    seo_url: string;
	
	    static createFrom(source: any = {}) {
	        return new SEOInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.seo_url = source["seo_url"];
	    }
	}
	export class TaxonomyPathItem {
	    id: string;
	    seo_info: SEOInfo;
	
	    static createFrom(source: any = {}) {
	        return new TaxonomyPathItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.seo_info = this.convertValues(source["seo_info"], SEOInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SEOVirtualCategory {
	    id: string;
	    taxonomy_path: TaxonomyPathItem[];
	
	    static createFrom(source: any = {}) {
	        return new SEOVirtualCategory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.taxonomy_path = this.convertValues(source["taxonomy_path"], TaxonomyPathItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MarketplaceListingRenderableTarget {
	    seo_virtual_category?: SEOVirtualCategory;
	
	    static createFrom(source: any = {}) {
	        return new MarketplaceListingRenderableTarget(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.seo_virtual_category = this.convertValues(source["seo_virtual_category"], SEOVirtualCategory);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PreRecordedVideo {
	    id: string;
	    playable_url: string;
	
	    static createFrom(source: any = {}) {
	        return new PreRecordedVideo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.playable_url = source["playable_url"];
	    }
	}
	export class Target {
	    id: string;
	    marketplace_listing_title: string;
	    listing_photos: ListingPhoto[];
	    pre_recorded_videos: PreRecordedVideo[];
	    listing_price: ListingPrice;
	    location_text: LocationText;
	    redacted_description: Description;
	    creation_time: number;
	    attribute_data: Attribute[];
	    commerce_badges_info: Commerce;
	
	    static createFrom(source: any = {}) {
	        return new Target(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.marketplace_listing_title = source["marketplace_listing_title"];
	        this.listing_photos = this.convertValues(source["listing_photos"], ListingPhoto);
	        this.pre_recorded_videos = this.convertValues(source["pre_recorded_videos"], PreRecordedVideo);
	        this.listing_price = this.convertValues(source["listing_price"], ListingPrice);
	        this.location_text = this.convertValues(source["location_text"], LocationText);
	        this.redacted_description = this.convertValues(source["redacted_description"], Description);
	        this.creation_time = source["creation_time"];
	        this.attribute_data = this.convertValues(source["attribute_data"], Attribute);
	        this.commerce_badges_info = this.convertValues(source["commerce_badges_info"], Commerce);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Root {
	    product_details_type: string;
	    target: Target;
	    marketplace_listing_renderable_target?: MarketplaceListingRenderableTarget;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.product_details_type = source["product_details_type"];
	        this.target = this.convertValues(source["target"], Target);
	        this.marketplace_listing_renderable_target = this.convertValues(source["marketplace_listing_renderable_target"], MarketplaceListingRenderableTarget);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	

}

