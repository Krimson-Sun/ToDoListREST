package main

import (
	todo "TodoListREST"
	"TodoListREST/pkg/handler"
	"TodoListREST/pkg/repository"
	"TodoListREST/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("Failed to load config:", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   viper.GetString("db.dbname"),
		SslMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatal("failed to initialize database: ", err)
	}

	log.Default().Print("Database connected")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
