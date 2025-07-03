package tools

import (
	log "github.com/sirupsen/logrus"
)

// LoginDetails Database collections
type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &mockDb{}
	var err error = db.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}
