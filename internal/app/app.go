package app

import (
	"os"

	"github.com/bank_service/internal/handler"
	"github.com/bank_service/internal/repository"
	"github.com/bank_service/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RunApp() {
	// logger
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// init config
	if err := InitConfig(); err != nil {
		logrus.Fatalf("cannot read config: %s", err.Error())
	}

	// load env variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
	}

	// init repository
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatal("failed to init DB: ", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	// init server
	server := new(Server)
	if err := server.Run(handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running server: %s", err.Error())
	}

	// graceful shutdown

}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
