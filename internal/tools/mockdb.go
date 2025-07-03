package tools

import (
	"time"
)

type mockDb struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "456DEF",
		Username:  "bob",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"bob": {
		Coins:    200,
		Username: "bob",
	},
}

// function of mockDb struct that implements the DatabaseInterface
func (d *mockDb) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDb) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDb) SetupDatabase() error {
	// No setup needed for mock database
	return nil
}
