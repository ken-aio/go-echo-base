package conf

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var (
	devDBUser     string = "postgres"
	devDBPassword string = ""
	devDBName     string = "xxx_dev"
	devDBHost     string = "localhost"
	devDBPort     string = "5432"

	testDBUser     string = "postgres"
	testDBPassword string = ""
	testDBName     string = "xxx_test"
	testDBHost     string = "localhost"
	testDBPort     string = "5432"

	prodDBUser     string = os.Getenv("DATABASE_USERNAME")
	prodDBPassword string = os.Getenv("DATABASE_PASSWORD")
	prodDBName     string = "xxx_production"
	prodDBHost     string = "localhost"
	prodDBPort     string = "5432"
)

func GetConfig() (c *Config) {
	if os.Getenv("ECHO_ENV") == "production" {
		c = &Config{
			DBUser:     prodDBUser,
			DBPassword: prodDBPassword,
			DBName:     prodDBName,
			DBHost:     prodDBHost,
			DBPort:     prodDBPort,
		}
	} else {
		c = &Config{
			DBUser:     devDBUser,
			DBPassword: devDBPassword,
			DBName:     devDBName,
			DBHost:     devDBHost,
			DBPort:     devDBPort,
		}
	}
	return
}

func GetTestConfig() (c *Config) {
	c = &Config{
		DBUser:     testDBUser,
		DBPassword: testDBPassword,
		DBName:     testDBName,
		DBHost:     testDBHost,
		DBPort:     testDBPort,
	}
	return
}
