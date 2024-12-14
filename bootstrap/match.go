package bootstrap

import (
	"github.com/yakob-abada/restapi/auth"
	"github.com/yakob-abada/restapi/handler"
	"github.com/yakob-abada/restapi/mapper"
	"github.com/yakob-abada/restapi/repo"
	"gorm.io/gorm"
)

func NewMatchHandler(db *gorm.DB) *handler.Match {
	return handler.NewMatch(auth.NewLoggedInUser(), repo.NewMatch(db), mapper.NewMatchMapper())
}
