package services

import (
	"github.com/manu-4216/web-golang/mvc/domain"
	"github.com/manu-4216/web-golang/mvc/utils"
)

// GetUser gets the user from the domain
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
