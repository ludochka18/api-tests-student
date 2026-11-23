package favorites

import (
	"testing"

	base "api-tests-template/tests"
)

func TestFavoriteAccounts(t *testing.T) {
	base.SetupSuite()

	loginFavoriteUsers(t)
}
