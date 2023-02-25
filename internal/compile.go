package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Compile(output string) error {

	if MainFile == "" {
		fmt.Println("\nI can't find the main file")
		return errors.New("1")
	}

	if !strings.Contains(output, ".exe") { /* Which system to compile for*/
		if runtime.GOOS == "windows" {
			fmt.Println("\nI can't compile a linux file on windows")
			return errors.New("1")
		}
	} else {
		if runtime.GOOS == "linux" {
			fmt.Println("\nI can't compile a windows file on linux")
			return errors.New("1")
		}
	}

	if runtime.GOOS == "linux" {
		_, err := exec.Command("garble", "-seed=random", "-literals", "-tiny", "build", "-o", output, MainFile).Output()
		if err != nil {
			fmt.Println("\nCompile error: " + err.Error())
			return err
		}

	} else if runtime.GOOS == "windows" {

		var err error
		_, err = exec.Command("cmd", "/C", "garble", "-seed=random", "-literals", "-tiny", "build", "-o", output, MainFile).Output()

		if err != nil {
			fmt.Println("\nCompile error: " + err.Error())
			return err
		}

	} else {
		fmt.Println("\nUnsupported OS")
		os.Exit(0)
	}

	return nil

}
