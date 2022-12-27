package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	ErrorColor   = "\033[1;31m%s\033[0m"
	SuccessColor = "\033[1;32m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
)

func PrintE(text string) {

	whitte := color.New(color.FgHiWhite).SprintFunc()
	fmt.Printf(ErrorColor, "DIFUSCATOR: ")
	fmt.Printf("%s", whitte(text))
	fmt.Printf(ErrorColor, "\n")
	os.Exit(0)

}

func LogS(text string) {
	whitte := color.New(color.FgHiWhite).SprintFunc()
	fmt.Printf("%s", whitte(text))
	fmt.Printf(SuccessColor, "           [ OK ]\n")

}

func LogE(text string) {
	whitte := color.New(color.FgHiWhite).SprintFunc()
	fmt.Printf("%s", whitte(text))
	fmt.Printf(ErrorColor, "           [ ERR ]\n")
}
