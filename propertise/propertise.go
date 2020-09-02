package propertise

import (
	"context"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	if e := gotenv.Load(".env"); e != nil {
		log.Fatal(e)
	}
}

var (
	Host    = os.Getenv("DB_HOST")
	DB      = os.Getenv("DB")
	UserC   = os.Getenv("USER_COLLECTION")
	ClientC = os.Getenv("CLIENT_COLLECTION")
	PersonC = os.Getenv("PERSON_COLLECTION")

	Port = os.Getenv("PORT")
	CTX  = context.Background()
)
