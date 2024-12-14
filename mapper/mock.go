package mapper

import (
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/restapi/model"
)

type LikeMock struct {
	mock.Mock
}

func (l *LikeMock) MapList(matches []*model.Match) []*Like {
	args := l.Called(matches)
	return args.Get(0).([]*Like)
}

func (l *LikeMock) Map(match *model.Match) *Like {
	args := l.Called(match)
	return args.Get(0).(*Like)
}

type MatchMock struct {
	mock.Mock
}

func (m *MatchMock) MapList(matches []*model.Match) []*Match {
	args := m.Called(matches)
	return args.Get(0).([]*Match)
}

func (m *MatchMock) Map(match *model.Match) *Match {
	args := m.Called(match)
	return args.Get(0).(*Match)
}
