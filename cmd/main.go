package main

import (
	"github.com/byroni/peafowl/internal/peafowl"
	"github.com/byroni/peafowl/pkg/boltdb"
	"github.com/byroni/peafowl/pkg/discord"
	"github.com/byroni/peafowl/ports"
)

func main() {
	// Ports
	var peafowlService ports.Peafowl
	var discordService ports.ChatService
	var database ports.Database

	// Adapters
  discordService = discord.New()
  database = boltdb.New()

  // load existing configuration
  database.Open()
  defer database.Close()

	peafowlService = peafowl.New(discordService, database)

	peafowlService.Publish("https://gfycat.com/weightycriminalfairfly-overwatch-blizzard-entertainment")
}
