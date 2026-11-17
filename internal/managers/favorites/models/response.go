package models

import advertisementModels "api-tests-template/internal/managers/advertisements/models"

// FavoriteList содержит список избранных объявлений.
type FavoriteList struct {
	Items      []FavoriteItem `json:"items"`
	Page       int            `json:"page"`
	TotalPages int            `json:"total_pages"`
}

// FavoriteItem содержит объявление и ID записи в избранном.
type FavoriteItem struct {
	FavoriteID    string                            `json:"favorite_id"`
	Advertisement advertisementModels.Advertisement `json:"advertisement"`
}
