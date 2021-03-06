package models

import "time"

type Customer struct {
	ID        string
	FristName string
	LastName  string
	Username  string
	Phones    []Phone
	Adresses  []Adress
	Products  []Product
	Email     string
	Gender    string
	Birthdate time.Time
	Password  string // should be hashed and validate password should be 8 symbols
	Status    string
}

type Phone struct {
	ID      string
	UserID  string
	Numbers []int64
	Code    string
}

type Adress struct {
	ID          string
	Country     string
	City        string
	District    string
	PostalCodes []int64
}

type Product struct {
	ID          string
	Name        string
	Types       Type
	Cost        int64
	OrderNumber int64
	Amount      int64
	Currency    string
	Rating      int64
}

type Type struct {
	ID   int64
	Name string
}
