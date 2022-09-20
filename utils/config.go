package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver            string
	DBSource            string
	ServerAddress       string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
func LoadConfig() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error Loading Config")
		os.Exit(1)
	}
	LoadEnvToStruct()
}
func LoadEnvToStruct() Config {
	str := "15m"
	dur_str, _ := time.ParseDuration(str)
	configs := Config{
		DBDriver:            "",
		DBSource:            "",
		ServerAddress:       "",
		TokenSymmetricKey:   "",
		AccessTokenDuration: dur_str,
	}
	return configs
}

func (conf *Config) fillConfig() {
	conf.DBDriver = os.Getenv("DB_DRIVER")
	conf.DBSource = os.Getenv("DB_SOURCE")
	conf.ServerAddress = os.Getenv("SERVER_ADDRESS")
	conf.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
	str := os.Getenv("ACCESS_TOKEN_DURATION")
	dur_str, _ := time.ParseDuration(str)
	conf.AccessTokenDuration = dur_str
}
