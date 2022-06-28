package transformer

import (
	models "user-service/model"
)

type User struct {
	Userid string `json:"Userid"`
	Name string `json:"Name"`
}

func DisplayUser(data models.User) User {
	return User{
		Userid: data.Userid,
		Name: data.Name,
	}
}

func DisplayAllUsers(data []models.User) []User {
	var dataResponse []User

	for _, item := range data {
		dataResponse = append(dataResponse, User{
			Userid: item.Userid,
			Name: item.Name,
		})
	}

	return dataResponse
}