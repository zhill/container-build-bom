package main

import (
	set "github.com/deckarep/golang-set"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Package struct {
	Name     string
	Version  string
	Location string
	Type     string
	Licenses []string
	Size     int
	Language string
}

func FindPip() ([]Package, error) {
	// This is intended for post-build, so assumes you have pip installed in the container itself.
	pips := make([]interface{}, 0)

	for _, name := range []string{"pip", "pip2", "pip3"} {
		cmd := exec.Command(name, "freeze", "--all")
		log.Println("Executing ", cmd.String())

		output, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		for _, line := range strings.Split(string(output), "\n") {
			pair := strings.Split(line, "==")
			if len(pair) == 2 {
				pips = append(pips, Package{
					Name:     pair[0],
					Version:  pair[1],
					Location: "",
					Language: "python",
					Type:     "pip",
					Size:     -1,
					Licenses: []string{},
				})
			}
		}
	}

	pipsSet := set.NewSetFromSlice(pips)
	s := pipsSet.ToSlice()
	return []Package(s), nil
}


func FindNpm() ([]Package, error) {
	// Find all package-lock.json files

	return []Package{}, nil
}

func analyze(c *cli.Context) error {
	var err error
	var result []Package

	pips, err := FindPip()
	if err != nil {
		return err
	}
	result = append(result, pips...)

	npms, err := FindNpm()
	if err != nil {
		return err
	}

	result = append(result, npms...)

	var d []byte
	d, err = json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(d))
	return nil
}

func main() {
	app := cli.App{
		Name:  "container-bom",
		Usage: "Constructs a single BOM json result for artfacts found in the local filesystem",
		Commands: []*cli.Command{
			{
				Name:    "analyze",
				Aliases: []string{"a"},
				Usage:   "Detect all supported artifact types and output result in json form to stdout",
				Flags:   []cli.Flag{},
				Action:  analyze,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
