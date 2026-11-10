package assertions

import (
	"api-tests-template/internal/managers/advertisements/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ContainsAdvertisementByTitle проверяет точное название среди результатов.
func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()
	for _, item := range items {
		if item.Title == title {
			return
		}
	}
	assert.Fail(t, "Объявление не найдено", "В результатах отсутствует название %q", title)
}
