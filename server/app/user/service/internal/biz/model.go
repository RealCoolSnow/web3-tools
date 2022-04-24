package biz

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile string `gorm:"unique"`
	Pass   string
	Name   string
	Age    int64
}

type UserReply struct {
	Name, Mobile string
	Age, ID      int64
}

type UserForToken struct {
	Mobile string
	Pass   string
	ID     int
}
