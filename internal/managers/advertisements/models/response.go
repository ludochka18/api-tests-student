package models

// AdvertisementsList содержит найденные объявления.
type AdvertisementsList struct {
	Items []Advertisement `json:"items"`
}

// Advertisement содержит название объявления.
type Advertisement struct {
	Title string `json:"title"`
}
