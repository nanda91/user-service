package routes

import (
	"user-service/config/boot"
	"user-service/delivery"
	"user-service/usecase"
)

var Time string

func ServiceRoute(api boot.ConfigHandler, us usecase.UseCaseInterface) {

	handler := &delivery.ServiceServer{
		UseCase: us,
	}

	myWebGroup := api.E.Group("/MyWeb")

	myWebGroup.GET("/DisplayAllUsers", handler.DisplayAllUsers)
	myWebGroup.POST("/DisplayUser", handler.DisplayUser)

	api.E.Run(api.Config.GetString(`app.host`) + ":" + api.Config.GetString( `app.port` ))
}
