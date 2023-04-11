/* build a marco polo cli tool that takes in the word marco and returns polo.
 * The tool should also take in a flag --times that will allow the user to specify how many times the tool should print polo.
 */
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/* a function that handles only the logic of Marco Polo
it takes a string and an int as arguments and returns a string
*/

func MarcoPolo(s string, times int) string {
	output := ""
	if s == "Marco" {
		for i := 0; i < times; i++ {
			output += "Polo "
		}
	} else {
		output = "Who are you?"
	}
	return output
}

// run the cli program and accept the arguments and flags
func main() {
	app := &cli.App{
		Name:  "Marco Polo Complex",
		Usage: "A simple cli tool that takes in the word marco and returns polo otherwise it returns 'who are you?'",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "times",
				Aliases: []string{"t"},
				Usage:   "number of times to print polo",
			},
		},
		Action: func(c *cli.Context) error {
			var times = c.Int("times")
			var name = c.Args().First()
			fmt.Println("Hello", name, "you want to print polo", times, "times")
			fmt.Println(MarcoPolo(c.Args().First(), c.Int("times")))
			return nil
		},
	}
	result := app.Run(os.Args)
	if result != nil {
		panic(result)
	}
}
