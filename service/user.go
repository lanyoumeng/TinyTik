package service

import (
	"TinyTik/common"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewServerService() *UserService {
	return &UserService{
		DB: common.GetDB(),
	}
}
