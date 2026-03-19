package internal

type GetLocations struct {
	Count    int               `json:"count"`
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []LocationResults `json:"results"`
}

type LocationResults struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}
