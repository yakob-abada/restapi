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

type Match struct {
	auth        auth.Authorization
	matchRepo   repo.IMatch
	matchMapper mapper.MatchMap
}

func NewMatch(auth auth.Authorization, matchRepo repo.IMatch, matchMapper mapper.MatchMap) *Match {
	return &Match{
		auth:        auth,
		matchRepo:   matchRepo,
		matchMapper: matchMapper,
	}
}

func (m *Match) WeMatched(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 10
	}

	result, err := m.matchRepo.MatchWithRecipient(
		m.auth.UserId(),
		[]int{model.MatchStatusMatched}, repo.NewPaginatedRequest(offset, page),
	)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, m.matchMapper.MapList(result))
}
