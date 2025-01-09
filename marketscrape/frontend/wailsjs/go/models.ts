export namespace main {
	
	export class MarketplaceListingResult {
	    data: {[key: string]: any};
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new MarketplaceListingResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.error = source["error"];
	    }
	}

}

