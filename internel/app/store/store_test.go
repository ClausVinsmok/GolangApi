package store_test

import (
	"os"
	"testing"
)

var (
	databaceURL string
)

func TestMain(m *testing.M) {
	databaceURL = os.Getenv("DATABASE_URL")
	if databaceURL == "" {
		databaceURL = "host=localhost dbname=restapi_test sslmode=disable password=1010"
	}

	os.Exit(m.Run())
}
