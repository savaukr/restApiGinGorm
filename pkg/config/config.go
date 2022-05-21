package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	HttpAddr string `default: "8080"`
	DBAddr string `default: "host=localhost user=sl password=1111 dbname=messages port=5432 sslmode=disable TimeZone=Asia/Shanghai"`

}

func NewConfig() (*Config, error) {
	var s Config
    err := envconfig.Process("RAGG", &s)
    if err != nil {
        return nil, err
    }
	return &s, nil
}