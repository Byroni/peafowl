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

	// Check if configuration already exists


  // Set configuration
  // TODO: This should probably be set in an exposed peafowlService method
  discordService.SetWebhookURL("https://discord.com/api/webhooks/870815898993295371/RL07y0NzWuZumQ4VWUgJ121WVJ6E9iiEb0t_ffGgrkgUGKe0RHtMGhC80Lr5Jp9NEw41")

	peafowlService = peafowl.New(discordService, database)
	peafowlService.Publish()
}
