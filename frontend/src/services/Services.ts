import {BACKEND_BASE_URL_API} from "../config";
import PasteService from "./PasteService";
import {ApiServiceConfig} from "./ApiService";

const config: ApiServiceConfig = {
    baseUrl: BACKEND_BASE_URL_API
}

const pasteService = new PasteService(config);

export {pasteService}