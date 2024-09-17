export namespace main {
	
	export class AxiosRequestConfig {
	    url?: string;
	    method?: string;
	    baseURL?: string;
	    headers?: {[key: string]: string};
	    auth?: {[key: string]: string};
	    params?: any;
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new AxiosRequestConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.method = source["method"];
	        this.baseURL = source["baseURL"];
	        this.headers = source["headers"];
	        this.auth = source["auth"];
	        this.params = source["params"];
	        this.data = source["data"];
	    }
	}
	export class AxiosResponse {
	    status: number;
	    statusText: string;
	    headers: {[key: string]: string};
	    config: AxiosRequestConfig;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new AxiosResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.statusText = source["statusText"];
	        this.headers = source["headers"];
	        this.config = this.convertValues(source["config"], AxiosRequestConfig);
	        this.data = source["data"];
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

