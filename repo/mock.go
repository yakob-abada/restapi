package repo

import (
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/restapi/model"
)

type MatchMock struct {
	mock.Mock
}

func (m *MatchMock) MatchWithRecipient(userId int, statuses []int, paginatedReq *PaginatedRequest) ([]*model.Match, error) {
	args := m.Called(userId, statuses, paginatedReq)
	return args.Get(0).([]*model.Match), args.Error(1)
}

type ProfileMock struct {
	mock.Mock
}

func (p *ProfileMock) Explore(userId int, paginatedReq *PaginatedRequest) ([]*model.Profile, error) {
	args := p.Called(userId, paginatedReq)
	return args.Get(0).([]*model.Profile), args.Error(1)
}
