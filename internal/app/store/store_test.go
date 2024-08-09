package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=max password=Nt[ybrf575 dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}
