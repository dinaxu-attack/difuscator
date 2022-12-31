package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dinaxu-attack/difuscator/app"
	"github.com/fatih/color"
)

func help() {
	fmt.Println("Usage: \n difuscator --file file.go/file.exe --compile (optional) --obfuscate (optional)")
}

var BuildFileName string

func main() {

	if len(os.Args) < 1 {
		help()
		os.Exit(0)
	}

	flag.Usage = func() {
		help()
	}

	file := flag.String("file", "", "")
	compile := flag.Bool("compile", false, "")
	obfuscate := flag.Bool("obfuscate", false, "")
	flag.Parse()

	if *file == "" {
		help()
		os.Exit(0)
	}

	launch(*file, *compile, *obfuscate)
}

func launch(file string, compile, obfuscate bool) {

	_, err := ioutil.ReadFile(file)

	if err != nil {
		app.PrintE("Can't find the file: " + file)
	}

	if obfuscate {
		fmt.Print("Obfuscating... ")

		err := app.Obfuscate(file)
		if err != nil {
			fmt.Printf(app.ErrorColor, "\t\t[ ERR ]\n")
		} else {
			fmt.Printf(app.SuccessColor, "\t\t[ OK ]\n")
		}
	}

	time.Sleep(2 * time.Second)
	if compile {
		fmt.Print("Compiling...")
		err, filename := app.Compile(file)

		BuildFileName = filename
		if err != nil {
			fmt.Printf(app.ErrorColor, "\t\t[ ERR ]\n")
			os.Exit(0)
		} else {
			fmt.Printf(app.SuccessColor, "\t\t[ OK ]\n")
		}
	}

	if BuildFileName == "" {
		BuildFileName = file
	}
	// .....................

	if !obfuscate && !compile {
		fmt.Print("Adding garbage...")
		err := app.Garbage(BuildFileName)
		if err != nil {
			fmt.Printf(app.ErrorColor, "\t[ ERR ]\n")
		} else {
			fmt.Printf(app.SuccessColor, "\t[ OK ]\n")
		}
		time.Sleep(500 * time.Millisecond)
		// ......................
	}

	if !obfuscate && !compile {
		fmt.Printf("Compressing...")
		cont, err := ioutil.ReadFile(BuildFileName)
		if err != nil {
			app.PrintE(err.Error())
		}
		btext := []byte(base64.StdEncoding.EncodeToString([]byte(cont)))
		compressed, err := app.Compress(btext)
		if err != nil {
			fmt.Printf(app.ErrorColor, "\t\t[ ERR ]\n")
		} else {
			fmt.Printf(app.SuccessColor, "\t\t[ OK ]\n")
		}
		fmt.Printf("Encrypting...")

		err = app.AES(compressed, BuildFileName)

		if err != nil {
			fmt.Printf(app.ErrorColor, "\t\t[ ERR ]\n")
		} else {
			fmt.Printf(app.SuccessColor, "\t\t[ OK ]\n")
		}
	}

	color.Green(app.Stat(BuildFileName))

	fmt.Println("Github: https://github.com/dinaxu-attack/difuscator")
	color.HiWhite("\nYour build: " + color.GreenString(BuildFileName))

}
