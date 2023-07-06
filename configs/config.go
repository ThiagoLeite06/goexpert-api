package configs

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebserverPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

var myConfig *conf

func LoadConfig(path string) (*conf, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&myConfig)
	if err != nil {
		panic(fmt.Errorf("erro ao decodificar as configurações do banco de dados: %s", err))
	}

	myConfig.TokenAuth = jwtauth.New("HS256", []byte(myConfig.JWTSecret), nil)

	return myConfig, err
}
