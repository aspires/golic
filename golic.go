package golic

import (
	"bytes"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
	"text/template"
	"time"
)

type options struct {
	Year      int    `short:"y" long:"year" description:"License year" value-name:"YEAR"`
	Copyright string `short:"c" long:"copyright" description:"Copyright name" value-name:"NAME"`
	URL       string `short:"u" long:"url" description:"URL" value-name:"URL"`
	Email     string `short:"e" long:"email" description:"E-Mail" value-name:"EMAIL"`
}

type general struct {
	Output string `short:"o" long:"output" description:"Output file" value-name:"FILE"`
	List   bool   `short:"l" long:"list" description:"List supported licenses"`
}

func listLicenses() {
	for key, _ := range licenses {
		fmt.Printf("- %s\n", key)
	}
}

func Command() {
	parser := flags.NewParser(nil, flags.HelpFlag)
	parser.Usage = "[OPTIONS] LICENSE"

	// License options
	opts := &options{}
	parser.AddGroup("License options", opts)

	// general options
	gen := &general{}
	parser.AddGroup("general options", gen)

	args, err := parser.Parse()

	// Show help message if user supply invalid options
	if err != nil {
		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	// List supported licenses
	if gen.List == true {
		fmt.Println("Supported licenses:")
		listLicenses()
		os.Exit(0)
	}

	// Ensure only one argument which supplied
	if len(args) > 1 {
		fmt.Printf("Error: Invalid arguments supplied: %s\n", args)
		os.Exit(0)
	}

	if len(args) == 1 {
		licTmpl, ok := licenses[args[0]]
		if ok == false {
			fmt.Printf("Error: License %q is not supported\n", args[0])
			os.Exit(0)
		}

		// Set default year
		if opts.Year == 0 {
			opts.Year = time.Now().Year()
		}

		tmpl := template.Must(template.New("License").Parse(strings.TrimPrefix(licTmpl, "\n")))

		var output bytes.Buffer

		err = tmpl.Execute(&output, opts)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		if gen.Output != "" {
			f, err := os.Create(gen.Output)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			output.WriteTo(f)
			fmt.Printf("License file %q successfully created!\n", gen.Output)
		} else {
			output.WriteTo(os.Stdout)
		}
		os.Exit(0)
	}

	parser.WriteHelp(os.Stderr)
	os.Exit(0)
}
