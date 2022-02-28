package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "Hello world",
		Usage: "Say hello",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello world, from de cli")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
