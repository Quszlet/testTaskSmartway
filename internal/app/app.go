package app

import (
	"log"
	"github.com/Quszlet/testTaskSmartway/internal/handler"
	"github.com/Quszlet/testTaskSmartway/internal/repository"
	"github.com/Quszlet/testTaskSmartway/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/configs")
	return viper.ReadInConfig()
}

func StartApp() {
	if err:= SetupConfig(); err != nil {
		log.Fatalf("Cannot read the config file: %s", err.Error())
	}

	dataBase, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLmode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("database initialization error: %s", err.Error())
	}

	repo := repository.NewRepository(dataBase)

	service := service.NewService(repo)

	handler := handler.NewHandeler(service)

	gin := gin.Default()
	
	handler.SetupRoutes(gin)

	gin.Run(":" + viper.GetString("port"))
}