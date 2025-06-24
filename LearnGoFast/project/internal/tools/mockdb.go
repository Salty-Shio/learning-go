package tools

import (
	"time"
)

type mockDb struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1000,
		Username: "alex",
	},
	"jason": {
		Coins:    2000,
		Username: "jason",
	},
	"marie": {
		Coins:    3000,
		Username: "marie",
	},
}

func (db *mockDb) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate a delay to mimic database access
	time.Sleep(time.Second * 1)
	
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDb) GetUserCoins(username string) *CoinDetails {
	// Simulate a delay to mimic database access
	time.Sleep(time.Second * 1)
	
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDb) SetupDatabase() error {
	return nil
}