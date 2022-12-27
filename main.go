package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	if *obfuscate == true && *compile != true {
		fmt.Println("Unable to obfuscate a non .go file! Use flag --compile --obfuscate")
		os.Exit(0)
	}

	if *file == "" {
		help()
		os.Exit(0)
	}

	launch(*file, *compile, *obfuscate)
}

func launch(file string, compile, obfuscate bool) {

	if obfuscate {
		fmt.Print("Obfuscating... ")

		err := app.Obfuscate()
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

	fmt.Print("Adding garbage...")
	err := app.Garbage(BuildFileName)
	if err != nil {
		fmt.Printf(app.ErrorColor, "\t[ ERR ]\n")
	} else {
		fmt.Printf(app.SuccessColor, "\t[ OK ]\n")
	}
	time.Sleep(500 * time.Millisecond)
	// ......................

	fmt.Printf("Encrypting...")
	cont, err := ioutil.ReadFile(BuildFileName)
	if err != nil {
		app.PrintE(err.Error())
	}

	btext := []byte(base64.StdEncoding.EncodeToString([]byte(cont)))

	err = app.AES(btext, BuildFileName)

	if err != nil {
		fmt.Printf(app.ErrorColor, "\t\t[ ERR ]\n")
	} else {
		fmt.Printf(app.SuccessColor, "\t\t[ OK ]\n")
	}

	app.Stat(BuildFileName)

	color.Green(strconv.Itoa(int(app.Old/1024.0)) + " MB -> " + strconv.Itoa(int(app.New/1024.0)) + " MB")

	fmt.Println("Github: https://github.com/dinaxu-attack/difuscator")
	color.HiWhite("\nYour build: " + color.GreenString(BuildFileName))

}
