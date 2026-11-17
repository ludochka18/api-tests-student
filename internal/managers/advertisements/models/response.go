package models

// AdvertisementsList содержит найденные объявления.
type AdvertisementsList struct {
	Items []Advertisement `json:"items"`
}

// Advertisement содержит данные объявления из ответа API.
type Advertisement struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	UserID      string  `json:"user_id"`
	Photos      []Photo `json:"photos"`
}

// Photo содержит данные фотографии объявления.
type Photo struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	SortOrder int    `json:"sort_order"`
}

// ErrorResponse содержит сообщение об ошибке API.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Details string `json:"details"`
}
