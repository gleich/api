package main

import (
	"github.com/gleich/api/pkg/db"
	"github.com/gleich/api/pkg/handle"
	"github.com/gleich/api/pkg/schema"
	"github.com/gleich/lumber"
)

func main() {
	lumber.Info("Started up server")
	db.Connect()
	s := schema.Init()
	handle.Run(s)
}
