export interface ApiServiceConfig {
    baseUrl: string;
}

export default class ApiService {

    constructor(protected config: ApiServiceConfig) {
    }


    protected apiUrl(url: string) {
        return `${this.config.baseUrl}/${url.replace(/^\//, '')}`
    }
}