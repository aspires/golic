package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const basePath = "templates"

var supported = map[string]string{
	"MIT": "MIT",
}

type Options struct {
	Year      int    `short:"y" long:"year" description:"License year" value-name:"YEAR"`
	Copyright string `short:"c" long:"copyright" description:"Copyright name" required:"true" value-name:"NAME"`
	URL       string `short:"u" long:"url" description:"URL" value-name:"URL"`
	Email     string `short:"e" long:"email" description:"E-Mail" value-name:"EMAIL"`
}

type General struct {
	Output string `short:"o" long:"output" description:"Output file" optional-value:"LICENSE" value-name:"FILE"`
	List   bool   `short:"l" long:"list" description:"List supported licenses"`
}

func main() {
	parser := flags.NewParser(nil, flags.HelpFlag)

	parser.Usage = "[OPTIONS] LICENSE"

	// License options
	opts := &Options{}
	parser.AddGroup("License Options", opts)

	// General options
	gen := &General{}
	parser.AddGroup("General Options", gen)

	args, err := parser.Parse()

	// Show help message if user supply invalid options
	if err != nil {
		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	// Ensure only one argument which supplied
	if len(args) != 1 {
		fmt.Printf("Error: Invalid arguments supplied: %s\n", args)
		os.Exit(0)
	}

	licFile, ok := supported[args[0]]
	if ok == false {
		fmt.Printf("Error: License %q is not supported\n", args[0])
		os.Exit(0)
	}

	// Set default year
	if opts.Year == 0 {
		opts.Year = time.Now().Year()
	}

	licPath := filepath.Join(basePath, fmt.Sprintf("%s.txt", licFile))

	tmpl := template.Must(template.ParseFiles(licPath))

	output := os.Stdout

	// TODO: write to file
	// TODO: list all supported licenses

	err = tmpl.Execute(output, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	os.Exit(0)
}
