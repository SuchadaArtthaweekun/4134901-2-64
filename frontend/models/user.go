package models

import "time"

type User struct {
	ID       uint
	FullName string
	Email    string
	CreateAt time.Time
	UpdataAt time.Time
}


// 1. post Add user
// 2. model
//   2.1 encode password
//   2.2 insert db (user,pass-encoded)