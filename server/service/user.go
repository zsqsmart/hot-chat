package service

import (
	"red-server/global"
	"red-server/model"
	"red-server/utils"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) Create(userDto *model.UserDto) (*model.User, error) {
	password := utils.BcryptHash(userDto.Password)
	user := &model.User{
		Password: password,
		Name:     userDto.Name,
		NickName: userDto.NickName,
		Email:    userDto.Email,
		Phone:    userDto.Phone,
		Gender:   userDto.Gender,
		Avatar:   userDto.Avatar,
		Birthday: userDto.Birthday,
	}
	return user, s.Insert(user)
}

func (s *UserService) Search(keyword string) model.Users {
	users := make(model.Users, 0)
	if keyword == "" {
		return users
	}
	result := s.db.Where("nick_name LIKE ? OR email = ?", keyword, keyword).Limit(10).Find(&users)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}
	return users
}
