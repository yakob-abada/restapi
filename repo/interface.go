package repo

import "github.com/yakob-abada/restapi/model"

type IProfile interface {
	Explore(userId int, paginatedReq *PaginatedRequest) ([]*model.Profile, error)
}

type IMatch interface {
	MatchWithRecipient(userId int, statuses []int, paginatedReq *PaginatedRequest) ([]*model.Match, error)
}
