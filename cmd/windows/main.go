package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ToffaKrtek/go-timetracker-integrator/internal/config"
)

func main() {
	config := config.GetConfig(os.Stdout, bufio.NewReader(os.Stdin))
	fmt.Println(config)
}
