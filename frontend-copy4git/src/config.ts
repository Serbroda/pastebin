
const { REACT_APP_API_BASE_URL } = process.env;
const { REACT_APP_API_BASE_URL_USE_WINDOW_LOCATION } = process.env;

function getBaseUrl() {
    if (REACT_APP_API_BASE_URL_USE_WINDOW_LOCATION === 'true') {
        return `${window.location.origin}`;
    }
    return REACT_APP_API_BASE_URL !== undefined && REACT_APP_API_BASE_URL !== null
        ? (REACT_APP_API_BASE_URL as string)
        : '';
}

const BACKEND_BASE_URL: string = getBaseUrl();
const BACKEND_BASE_URL_API: string = BACKEND_BASE_URL + '/api';

export { BACKEND_BASE_URL, BACKEND_BASE_URL_API }