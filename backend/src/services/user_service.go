package services

import "organization-management-justin/backend/src/models"

func GetAllUsers() []models.User {
	return []models.User{
		{ID: 1, Name: "Justin", Email: "justin@mail.com"},
		{ID: 2, Name: "Andi", Email: "andi@mail.com"},
	}
}
