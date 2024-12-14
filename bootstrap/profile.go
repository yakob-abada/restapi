package bootstrap

import (
	"github.com/yakob-abada/restapi/auth"
	"github.com/yakob-abada/restapi/handler"
	"github.com/yakob-abada/restapi/repo"
	"gorm.io/gorm"
)

func NewProfileHandler(db *gorm.DB) *handler.Profile {
	return handler.NewProfile(auth.NewLoggedInUser(), repo.NewProfile(db))
}
