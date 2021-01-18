package main

import (
	"github.com/Matt-Gleich/api/pkg/handle"
	"github.com/Matt-Gleich/api/pkg/schema"
)

func main() {
	s := schema.Init()
	handle.Run(s)
}
