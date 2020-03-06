package main

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func FindPip() ([]string, error) {
	// Run pip freeze
	return []string{"pip1"}, nil
}

func FindNpm() ([]string, error) {
	// Find all package-lock.json files
	return []string{"npm1"}, nil
}

func analyze(c *cli.Context) error {
	var err error
	result := map[string][]string {
		"pip": nil,
		"npm": nil,
	}

	result["pip"], err = FindPip()
	if err != nil {
		return err
	}

	result["npm"], err = FindNpm()
	if err != nil {
		return err
	}

	var d []byte
	d, err = json.Marshal(result)
	log.Println("Results: ", string(d))
	return nil
}


func main() {
	app := cli.App{
		Name: "container-bom",
		Usage: "Constructs a single BOM json result for artfacts found in the local filesystem",
		Commands: []*cli.Command {
			{
				Name: "analyze",
				Aliases: []string{"a"},
				Usage: "Detect all supported artifact types and output result in json form to stdout",
				Flags: []cli.Flag{},
				Action: analyze,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
