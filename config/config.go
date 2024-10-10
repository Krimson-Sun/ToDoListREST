package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postrgres DataBase `yaml:"dbconfig"`
}

type DataBase struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
	SslMode  string `yaml:"sslmode"`
}

func LoadConfig(filename string) (*DataBase, error) {
	var cfg Config
	err := cleanenv.ReadConfig(filename, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Postrgres, nil

}

func (dbconfig *DataBase) GetDSN() string {
	fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbconfig.Host, dbconfig.User, dbconfig.Password, dbconfig.DBName, dbconfig.Port, dbconfig.SslMode)
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbconfig.Host, dbconfig.User, dbconfig.Password, dbconfig.DBName, dbconfig.Port, dbconfig.SslMode)
}
