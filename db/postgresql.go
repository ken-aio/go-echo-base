package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/ken-aio/pictwitter-go/conf"
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func Init() *dbr.Session {
	return getSession()
}

func InitTest() *dbr.Session {
	return getTestSession()
}

func getSession() *dbr.Session {
	c := conf.GetConfig()
	con, err := dbr.Open("postgres", "postgres://"+c.DBUser+":"+c.DBPassword+"@"+c.DBHost+":"+c.DBPort+"/"+c.DBName+"?sslmode=disable", nil)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return handleSession(con)
}

func getTestSession() *dbr.Session {
	c := conf.GetTestConfig()
	con, err := dbr.Open("postgres", "postgres://"+c.DBUser+":"+c.DBPassword+"@"+c.DBHost+":"+c.DBPort+"/"+c.DBName+"?sslmode=disable", nil)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return handleSession(con)
}

func handleSession(con *dbr.Connection) *dbr.Session {
	session := con.NewSession(nil)
	return session
}
