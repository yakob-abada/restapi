package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/restapi/auth"
	"github.com/yakob-abada/restapi/mapper"
	"github.com/yakob-abada/restapi/model"
	"github.com/yakob-abada/restapi/repo"
	"net/http"
	"strconv"
)

type Like struct {
	auth       auth.Authorization
	matchRepo  repo.IMatch
	likeMapper mapper.LikeMap
}

func NewLike(auth auth.Authorization, matchRepo repo.IMatch, likeMapper mapper.LikeMap) *Like {
	return &Like{
		auth:       auth,
		matchRepo:  matchRepo,
		likeMapper: likeMapper,
	}
}

func (l *Like) WhoLikedMe(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 10
	}

	result, err := l.matchRepo.MatchWithRecipient(
		l.auth.UserId(),
		[]int{model.MatchStatusPending}, repo.NewPaginatedRequest(offset, page),
	)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, l.likeMapper.MapList(result))
}
