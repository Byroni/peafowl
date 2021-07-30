package main

import (
	"github.com/byroni/peafowl/gui"
	"github.com/byroni/peafowl/internal/peafowl"
	"github.com/byroni/peafowl/ports"
)

func main() {
	var peafowlService ports.Peafowl

	peafowlService = peafowl.New()

	// TUI initialization
	gui.New(peafowlService)
}
