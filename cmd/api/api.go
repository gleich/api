package main

import (
	"github.com/gleich/api/pkg/db"
	"github.com/gleich/api/pkg/handle"
	"github.com/gleich/api/pkg/schema"
	"github.com/gleich/lumber"
)

func main() {
	lumber.Info("Started up server")
	err := db.Connect()
	if err != nil {
		lumber.Fatal(err, "Failed to connect to database")
	}
	s := schema.Init()
	handle.Run(s)
}
