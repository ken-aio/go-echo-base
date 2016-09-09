package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/ken-aio/go-echo-base/conf"
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
	con, err := dbr.Open("postgres", "postgres://"+conf.DB_USER+":"+conf.DB_PASSWORD+"@"+conf.DB_HOST+":"+conf.DB_PORT+"/"+conf.DB_NAME+"?sslmode=disable", nil)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return handleSession(con)
}

func getTestSession() *dbr.Session {
	con, err := dbr.Open("postgres", "postgres://"+conf.DB_USER+":"+conf.DB_PASSWORD+"@"+conf.DB_HOST+":"+conf.DB_PORT+"/"+conf.TEST_DB_NAME+"?sslmode=disable", nil)

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
