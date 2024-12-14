package bootstrap

import (
	"github.com/yakob-abada/restapi/auth"
	"github.com/yakob-abada/restapi/handler"
	"github.com/yakob-abada/restapi/mapper"
	"github.com/yakob-abada/restapi/repo"
	"gorm.io/gorm"
)

func NewLikeHandler(db *gorm.DB) *handler.Like {
	return handler.NewLike(auth.NewLoggedInUser(), repo.NewMatch(db), mapper.NewLikeMapper())
}
