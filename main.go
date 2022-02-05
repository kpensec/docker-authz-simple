package main

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"log"
	"fmt"
)

func main() {
	fmt.Printf("Starting authorization plugin")

	plugin, err := newPlugin()
	if err != nil {
		log.Fatal(err)
	}

	h := authorization.NewHandler(plugin)

	if err := h.ServeUnix("authz-plugin", 0); err != nil {
		log.Fatal(err)
	}
}
