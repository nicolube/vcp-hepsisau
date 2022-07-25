package model

type UserModel struct {
	Model
	FirstName   string
	LastName    string
	Email       string
	Password    string
	LastLoginAt UnixTimeS
	SingedUpAt  UnixTimeS
}
