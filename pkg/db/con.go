package db

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gleich/lumber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dbURL, err := url.Parse(os.Getenv("DB_URL"))
	if err != nil {
		lumber.Fatal(err, "Failed to")
	}
	password, _ := dbURL.User.Password()

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbURL.Hostname(),
		dbURL.User.Username(),
		password,
		strings.TrimPrefix(dbURL.Path, "/"),
		dbURL.Port(),
	)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	lumber.Success("Connected to database")
	return nil
}
