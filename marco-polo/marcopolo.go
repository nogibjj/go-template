// build a marco polo cli tool that takes in the word marco and returns polo otherwise it returns "I don't understand"
package main

// imports as package "cli"
import (
	"os"

	"github.com/urfave/cli/v2"
)

func MarcoPolo(s string) string {
	if s == "Marco" {
		return "Polo"
	} else {
		return "Huh?"
	}
}

// run the program using cli package
func main() {
	app := &cli.App{
		Name:  "Marco Polo",
		Usage: "A simple cli tool that takes in the word marco and returns polo otherwise it returns huh?",
		Action: func(c *cli.Context) error {
			println(MarcoPolo(c.Args().First()))
			return nil
		},
	}
	//capture the result of the cli app
	result := app.Run(os.Args)
	if result != nil {
		panic(result)
	}
}
