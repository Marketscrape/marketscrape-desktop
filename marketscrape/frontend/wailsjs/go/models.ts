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
	export class Details {
	    format: string;
	    family: string;
	    families: any;
	    parameter_size: string;
	    quantization_level: string;
	
	    static createFrom(source: any = {}) {
	        return new Details(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.format = source["format"];
	        this.family = source["family"];
	        this.families = source["families"];
	        this.parameter_size = source["parameter_size"];
	        this.quantization_level = source["quantization_level"];
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
	export class Model {
	    name: string;
	    modified_at: string;
	    size: number;
	    digest: string;
	    details: Details;
	
	    static createFrom(source: any = {}) {
	        return new Model(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.modified_at = source["modified_at"];
	        this.size = source["size"];
	        this.digest = source["digest"];
	        this.details = this.convertValues(source["details"], Details);
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
	export class VehicleSpecifications {
	    co2_emissions: any;
	    engine_size: any;
	    gas_mileage_city: any;
	    gas_mileage_combined: any;
	    gas_mileage_highway: any;
	    horse_power: any;
	    safety_rating_front: any;
	    safety_rating_overall: any;
	    safety_rating_rollover: any;
	    safety_rating_side: any;
	    safety_rating_side_barrier: any;
	
	    static createFrom(source: any = {}) {
	        return new VehicleSpecifications(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.co2_emissions = source["co2_emissions"];
	        this.engine_size = source["engine_size"];
	        this.gas_mileage_city = source["gas_mileage_city"];
	        this.gas_mileage_combined = source["gas_mileage_combined"];
	        this.gas_mileage_highway = source["gas_mileage_highway"];
	        this.horse_power = source["horse_power"];
	        this.safety_rating_front = source["safety_rating_front"];
	        this.safety_rating_overall = source["safety_rating_overall"];
	        this.safety_rating_rollover = source["safety_rating_rollover"];
	        this.safety_rating_side = source["safety_rating_side"];
	        this.safety_rating_side_barrier = source["safety_rating_side_barrier"];
	    }
	}
	export class VehicleOdometerData {
	    unit: string;
	    value: number;
	
	    static createFrom(source: any = {}) {
	        return new VehicleOdometerData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.unit = source["unit"];
	        this.value = source["value"];
	    }
	}
	export class VehicleData {
	    vehicle_carfax_report: any;
	    vehicle_condition: any;
	    vehicle_exterior_color: string;
	    vehicle_features: any[];
	    vehicle_fuel_type: any;
	    vehicle_identification_number: any;
	    vehicle_interior_color: string;
	    vehicle_is_paid_off: any;
	    vehicle_make_display_name: string;
	    vehicle_model_display_name: string;
	    vehicle_number_of_owners: any;
	    vehicle_odometer_data: VehicleOdometerData;
	    vehicle_registration_plate_information: any;
	    vehicle_seller_type: string;
	    vehicle_specifications: VehicleSpecifications;
	    vehicle_title_status: any;
	    vehicle_transmission_type: string;
	    vehicle_trim_display_name: any;
	    vehicle_website_link: any;
	
	    static createFrom(source: any = {}) {
	        return new VehicleData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.vehicle_carfax_report = source["vehicle_carfax_report"];
	        this.vehicle_condition = source["vehicle_condition"];
	        this.vehicle_exterior_color = source["vehicle_exterior_color"];
	        this.vehicle_features = source["vehicle_features"];
	        this.vehicle_fuel_type = source["vehicle_fuel_type"];
	        this.vehicle_identification_number = source["vehicle_identification_number"];
	        this.vehicle_interior_color = source["vehicle_interior_color"];
	        this.vehicle_is_paid_off = source["vehicle_is_paid_off"];
	        this.vehicle_make_display_name = source["vehicle_make_display_name"];
	        this.vehicle_model_display_name = source["vehicle_model_display_name"];
	        this.vehicle_number_of_owners = source["vehicle_number_of_owners"];
	        this.vehicle_odometer_data = this.convertValues(source["vehicle_odometer_data"], VehicleOdometerData);
	        this.vehicle_registration_plate_information = source["vehicle_registration_plate_information"];
	        this.vehicle_seller_type = source["vehicle_seller_type"];
	        this.vehicle_specifications = this.convertValues(source["vehicle_specifications"], VehicleSpecifications);
	        this.vehicle_title_status = source["vehicle_title_status"];
	        this.vehicle_transmission_type = source["vehicle_transmission_type"];
	        this.vehicle_trim_display_name = source["vehicle_trim_display_name"];
	        this.vehicle_website_link = source["vehicle_website_link"];
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
	    vehicle_data?: VehicleData;
	
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
	        this.vehicle_data = this.convertValues(source["vehicle_data"], VehicleData);
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

