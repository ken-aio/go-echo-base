package api

import (
	"database/sql"
	"log"
	"testing"
	_ "github.com/lib/pq"
	. "gopkg.in/check.v1"
	"gopkg.in/testfixtures.v1"
)

const FixturesPath = "../fixtures"

func TestPackage(t *testing.T) {
	con, err := sql.Open("postgres", "dbname=xxx_test  sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	testfixtures.SkipDatabaseNameCheck(true)
	err = testfixtures.LoadFixtures(FixturesPath, con, &testfixtures.PostgreSQLHelper{})
	if err != nil {
		log.Fatal(err)
	}
	TestingT(t)
}

var _ = Suite(&ApiSuite{})
