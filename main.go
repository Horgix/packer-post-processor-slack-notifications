package main

import (
	"log"

	"github.com/horgix/packer-post-processor-slack-notifications/slack-notifications"

	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	log.Println("Starting...")
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterPostProcessor(new(slacknotifications.PostProcessor))
	server.Serve()
}
