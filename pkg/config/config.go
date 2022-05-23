package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	"os"
	"log"
)

type Config struct {
	HttpAddr string `default: ":8080"`
	DBAddr string `default: "host=localhost user=postgres password= dbname=messages port=5432 sslmode=disable"`

}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
	db_username, exists_username :=os.LookupEnv("DB_USERNAME")
	db_userpassword, exists_userpassword :=os.LookupEnv("DB_USER_PASSWORD")
	db_port, exists_dbport :=os.LookupEnv("DB_PORT")
	db_name, exists_dbname :=os.LookupEnv("DB_NAME")



	var s Config
	s.HttpAddr = ":8081"
	if (exists_username && exists_dbname && exists_userpassword && exists_dbport) {
		s.DBAddr = "host=localhost user=" + db_username + " password=" + db_userpassword +" dbname="+db_name+ " port="+db_port+ " sslmode=disable"
	}

    err := envconfig.Process("RAGG", &s)
    if err != nil {
        return nil, err
    }
	return &s, nil
}