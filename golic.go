// An open source license generator for your projects
package main

import (
	"bytes"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/subosito/golic/templates"
	"os"
	"path"
	"text/template"
	"time"
)

func main() {
	parser := flags.NewNamedParser(path.Base(os.Args[0]), flags.None)
	parser.Usage = "[OPTIONS] LICENSE"

	// Help options
	var help struct {
		ShowHelp bool `short:"h" long:"help" description:"Show this help message"`
	}
	parser.AddGroup("Help Options", &help)

	// License options
	var options struct {
		Year      int    `short:"y" long:"year" description:"License year" value-name:"YEAR"`
		Copyright string `short:"c" long:"copyright" description:"Copyright name" value-name:"NAME"`
		URL       string `short:"u" long:"url" description:"URL" value-name:"URL"`
		Email     string `short:"e" long:"email" description:"E-Mail" value-name:"EMAIL"`
	}

	parser.AddGroup("License Options", &options)

	// general options
	var general struct {
		Output string `short:"o" long:"output" description:"Output file" value-name:"FILE"`
		List   bool   `short:"l" long:"list" description:"List supported licenses"`
	}

	parser.AddGroup("General Options", &general)

	args, err := parser.Parse()

	// Show help message if user supply invalid options
	if err != nil {
		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	// List supported licenses
	if general.List == true {
		fmt.Println("Supported licenses:")
		for _, val := range templates.List() {
			fmt.Printf("- %s\n", val)
		}
		os.Exit(0)
	}

	// Ensure only one argument which supplied
	if len(args) > 1 {
		fmt.Printf("Error: Invalid arguments supplied: %s\n", args)
		os.Exit(0)
	}

	if len(args) == 1 {
		lic, ok := templates.Load(args[0])
		if ok == false {
			fmt.Printf("Error: License %q is not supported\n", args[0])
			os.Exit(0)
		}

		// Set default year
		if options.Year == 0 {
			options.Year = time.Now().Year()
		}

		tmpl := template.Must(templates.Template(lic.Template))

		var output bytes.Buffer

		err = tmpl.Execute(&output, options)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		if general.Output != "" {
			f, err := os.Create(general.Output)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			output.WriteTo(f)
			fmt.Printf("License file %q successfully created!\n", general.Output)
		} else {
			output.WriteTo(os.Stdout)
		}
		os.Exit(0)
	}

	parser.WriteHelp(os.Stderr)
	os.Exit(0)
}
