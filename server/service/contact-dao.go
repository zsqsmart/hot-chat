package service

import (
	"red-server/global"
	"red-server/model"
	"red-server/utils"
)

func (s *ContactService) Insert(contact *model.Contact) error {
	result := s.db.Create(contact)
	return result.Error
}

func (s *ContactService) IsMyContact(targetId int64, contactType model.ContactType) bool {
	contact := model.Contact{}
	currentUser := utils.GetCurrentUser(s.ctx)
	result := s.db.Where("owner_id = ? AND target_id = ? AND type = ?", currentUser.Id, targetId, contactType).First(&contact)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}
	return contact.Id != 0
}

func (s *ContactService) IsMyFriend(friendId int64) bool {
	return s.IsMyContact(friendId, model.ContactTypeFriend)
}

func (s *ContactService) IsMyCommunity(communityId int64) bool {
	return s.IsMyContact(communityId, model.ContactTypeCommunity)
}

// 互相的朋友
func (s *ContactService) IsFriend(userId1, userId2 int64) bool {
	contact := model.Contact{}
	result := s.db.Where("(owner_id = ? AND target_id = ?) OR (owner_id = ? AND target_id = ?) AND type = ?", userId1, userId2, userId2, userId1, model.ContactTypeFriend).First(&contact)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}
	return contact.Id != 0
}

func (s *ContactService) GetAllByOwnerId(ownerId int64) []*model.Contact {
	contacts := []*model.Contact{}
	result := s.db.Where("owner_id = ?", ownerId).Find(&contacts)
	if result.Error != nil {
		global.Logger.Error(result.Error)
	}
	return contacts
}
