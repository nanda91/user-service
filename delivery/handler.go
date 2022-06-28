package delivery

import (
	"github.com/gin-gonic/gin"
	httpResponse "user-service/delivery/response"
	"user-service/delivery/transformer"
	"user-service/usecase"
)

type ServiceServer struct {
	UseCase  usecase.UseCaseInterface
	Response httpResponse.HTTPResponseHelper
}

type InputFindUser struct {
	Userid string `json:"Userid" binding:"required"`
}

func  (s *ServiceServer) DisplayAllUsers(c *gin.Context) {
	data, err := s.UseCase.FindAllUser()
	if err != nil {
		s.Response.SendBadRequest(c, "Failed get data users")
		return
	}

	dataResponse := transformer.DisplayAllUsers(data)
	s.Response.SendSuccess(c, dataResponse)
}

func  (s *ServiceServer) DisplayUser(c *gin.Context) {
	var input InputFindUser

	if err := c.ShouldBindJSON(&input); err != nil {
		s.Response.SendBadRequest(c, err.Error())
		return
	}

	data, err := s.UseCase.FindUser(input.Userid)
	if err != nil {
		s.Response.SendBadRequest(c, "Failed get data user")
		return
	}

	dataResponse := transformer.DisplayUser(data)
	s.Response.SendSuccess(c, dataResponse)
}