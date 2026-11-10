package advertisements

import (
	client "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

// GetAdvertisementsList возвращает результаты поиска по названию.
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	t.Helper()
	raw := client.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err)
	var result models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &result))
	return result
}
