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
	junk := flag.Bool("j", false, "")
	garbage := flag.Bool("g", false, "")
	compile := flag.Bool("c", false, "")
	flag.Parse()

	if *compile && *output == "" {
		help()
	}

	if !*obfuscate && !*compile {
		help()
	}

	launch(*compile, *obfuscate, *output, *junk, *garbage)
}

func launch(compile bool, obf bool, output string, junk bool, garbage bool) {

	go func() {
		for per := 0; per <= 100; per++ {
			if per == 99 {
				per = 0
				continue
			}
			Progressbar(per)
		}
	}()

	internal.CopyFiles("./")
	if obf {
		internal.ObfuscateFiles("./" + internal.TempName)
	}
	if compile {

		internal.GetMain("./" + internal.TempName)

		err := internal.Compile(output)

		if err != nil {
			os.Exit(0)
		}

	}
	if garbage {
		internal.Garbage(output)
	}

	if junk {
		internal.GetMain("./" + internal.TempName)
		internal.JunkCode()
	}

	fmt.Printf("\r[##################################################]")
	println()
	color.Green(internal.Stat(output))
}

func help() {

	fmt.Println(`Usage: difuscator [-o (Example build.exe) ] [-c (bool) ] [-obf (bool) ] [-j (bool) ] [-g (bool) ]` + "\n\n-o: Output. Build name\n-c: Compile\n-obf: Obfuscation (Random files name)\n-j: Junk code (Random go code)\n-g: Garbage")
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
