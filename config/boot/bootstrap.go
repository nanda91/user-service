package boot

import (
	"github.com/gin-gonic/gin"
	"log"
	cfg "user-service/config/env"
	"user-service/lib/core/postgresql"
	"user-service/repository"
	"user-service/usecase"
)


type ConfigHandler struct {
	E *gin.Engine
	Config cfg.Config
}

var (
	db 	postgresql.Info
	err error
)

func (h *ConfigHandler) RegisterConfig() usecase.UseCaseInterface {
	db = postgresql.Info{
		Hostname: h.Config.GetString("database.host"),
		Database: h.Config.GetString("database.database"),
		Username: h.Config.GetString("database.username"),
		Password: h.Config.GetString("database.password"),
	}

	dbConnect, err := db.Connect()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	repo 	:= repository.NewLogRepository(dbConnect, h.Config)
	useCase := usecase.NewUseCase(h.Config, repo)

	return useCase
}