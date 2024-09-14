package main

import (
	"bufio"
	"os"

	"github.com/ToffaKrtek/go-timetracker-integrator/internal/config"
	"github.com/ToffaKrtek/go-timetracker-integrator/internal/tracker"
)

func main() {
	config := config.GetConfig(os.Stdout, bufio.NewReader(os.Stdin))
  tracker.Run(config)
}
