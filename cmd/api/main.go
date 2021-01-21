package main

import (
	"github.com/Matt-Gleich/api/pkg/handle"
	"github.com/Matt-Gleich/api/pkg/schema"
	"github.com/Matt-Gleich/lumber"
)

func main() {
	lumber.Info("Started up server")
	s := schema.Init()
	handle.Run(s)
}
