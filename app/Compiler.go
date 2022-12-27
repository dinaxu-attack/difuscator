package app

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func Compile(file string) (error, string) {
	d := strconv.Itoa(int(Random(100, 10000)))
	var buildName string
	if runtime.GOOS == "linux" {
		buildName = "build_" + d
		_, err := exec.Command("garble", "-seed=random", "-literals", "-tiny", "build", "-o", buildName, file).Output()
		if err != nil {
			return err, ""
		}

	} else if runtime.GOOS == "windows" {
		buildName = "build_" + d + ".exe"
		_, err := exec.Command("cmd", "/C", "garble", "-seed=random", "-literals", "-tiny", "build", "-o", buildName, file).Output()
		if err != nil {
			return err, ""
		}

	} else {
		PrintE("Unsupported OS")
		os.Exit(0)
	}

	return nil, buildName

}
