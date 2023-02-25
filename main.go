package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dinaxu-attack/difuscator/internal"

	"github.com/fatih/color"
)

func main() {

	defer timer()()

	if len(os.Args) < 1 {
		help()

	}

	flag.Usage = func() {
		help()

	}

	output := flag.String("o", "", "")
	obfuscate := flag.Bool("obf", false, "")
	compile := flag.Bool("c", false, "")
	flag.Parse()

	if *compile && *output == "" {
		help()
	}

	if !*obfuscate && !*compile {
		help()
	}

	launch(*compile, *obfuscate, *output)
}

func launch(compile bool, obf bool, output string) {

	go func() {
		for per := 0; per <= 100; per++ {
			if per == 99 {
				per = 0
				continue
			}
			Progressbar(per)
		}
	}()

	if obf {
		internal.ObfuscateFiles("./")
	}
	if compile {

		internal.GetMain("./")

		err := internal.Compile(output)

		if err != nil {
			os.Exit(0)
		}

		fmt.Printf("\r[##################################################]")
		internal.Garbage(output)
		println()
		color.Green(internal.Stat(output))
	}
}

func help() {

	fmt.Println(`Usage: difuscator [-o output (Example build.exe) ] [-c compile (bool) ] [-obf (bool) ]`)
	os.Exit(0)
}

func Progressbar(per int) {

	chars := per * 50 / 100

	fmt.Printf("\r[")

	for char := 0; char < chars; char++ {
		fmt.Printf("#")
	}
	for charx := 0; charx < 50-chars; charx++ {
		fmt.Printf(" ")
	}
	fmt.Printf("]")
	time.Sleep(10 * time.Millisecond)
}

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Execution time: %v\n", time.Since(start))
	}
}
