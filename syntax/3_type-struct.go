package main

import (
	"log"
	"time"
)

// type structで自作の型を作成（swiftと似てる）
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Toshiki",
		LastName:  "Yanagimoto",
	}

	log.Println(user.FirstName, user.LastName)
	log.Println(user.PhoneNumber) // 初期値：""
	log.Println(user.BirthDate)   // 初期値：0001-01-01 00:00:00 +0000 UTC
}
