package test

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Name string
}

// `Profile` 属于 `User`， `UserID` 是外键
type Profile struct {
	gorm.Model
	UserID int
	User   UserInfo
	Name   string
}
