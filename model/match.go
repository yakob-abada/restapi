package model

import (
	"time"
)

const (
	MatchStatusPending = iota
	MatchStatusMatched
	MatchStatusUnMatched
)

type Match struct {
	RecipientUserId int  `gorm:"primaryKey;autoIncrement:false"`
	ActorUserId     int  `gorm:"primaryKey;autoIncrement:false"`
	Status          int8 `json:"-"`
	CreatedAt       time.Time
}
