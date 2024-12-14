package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/restapi/auth"
	"github.com/yakob-abada/restapi/repo"
	"net/http"
	"strconv"
)

type Profile struct {
	profileRep repo.IProfile
	auth       auth.Authorization
}

func NewProfile(auth auth.Authorization, profileRep repo.IProfile) *Profile {
	return &Profile{
		auth:       auth,
		profileRep: profileRep,
	}
}

func (p *Profile) Explore(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 10
	}

	list, err := p.profileRep.Explore(p.auth.UserId(), repo.NewPaginatedRequest(offset, page))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, list)
}
