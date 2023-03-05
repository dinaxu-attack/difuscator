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

	var typee int
	if MainFile == "" {
		fmt.Println("\nI can't find the main file")
		return errors.New("1")
	}

	if !strings.Contains(output, ".exe") { /* Which system to compile for*/
		typee = 0
		if runtime.GOOS == "windows" {
			fmt.Println("\nI can't compile a linux file on windows")
			return errors.New("1")
		}
	} else {
		typee = 1
	}

	if runtime.GOOS == "linux" {
		cmd := exec.Command("env")
		var out strings.Builder
		env := os.Environ()

		if typee == 1 {
			env = append(env, "GOOS=windows")
		} else {
			env = append(env, "GOOS=linux")
		}
		cmd.Env = env
		cmd.Stdout = &out
		cmd.Run()

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
