package core

import "time"

var supportedCurrencies []string = []string{"USD", "RUB", "NGN"}

type User struct {
	ID         int
	Name       string
	Accounts   []Account
	Created_at time.Time
}

type Account struct {
	ID         int
	Owner      *User
	Balance    float32
	Currency   string
	Created_at time.Time
}

type Error struct {
	Message string
}
