package internal

import (
	_ "embed"
	"errors"
	"fmt"
	"io/ioutil"
	random "math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	vartypes  = []string{"string", "int", "bool", "byte"}
	boolvalue = []string{"true", "false"}
	funcs     = make([]string, 0)
)

/*
Generates junk code (Random go variables, funcs)
Like var dsjkjasdjksadjljlk = "ASJKLDjkladsjkla"
or func DHdhjhudashaidh() {}
*/
func JunkCode() {

	for i := 0; i <= int(Random(5, 10)); i++ {
		dir := TempName + RandomString(5, 10)
		os.Mkdir(dir, os.ModePerm)

		for i := 0; i < int(Random(4, 10)); i++ {
			file := dir + "/" + RandomString(7, 10) + ".go"

			f, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			defer f.Close()

			f.WriteString(string("package " + strings.Split(dir, TempName)[1] + "\n\n"))
			for i := 0; i < int(Random(2, 5)); i++ {
				f.WriteString(string(randomFunc(strings.Split(dir, TempName)[1])))
			}
		}

		err := AddToMain(strings.Split(dir, TempName)[1])

		if err != nil {
			return
		}
	}

	AppendMain()

}

// Creates a random function and in that function calls others
func AppendMain() {
	f, _ := os.OpenFile(MainFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	f.WriteString("\n\n func " + UniqueName() + "() {")

	for _, funct := range funcs {
		f.WriteString("\n    " + funct + "()\n")
	}

	f.WriteString("\n}")
}

var gomod []byte

func AddToMain(folder string) error {
	if MainFile == "" {
		fmt.Println("\nI can't find the main file")
		return errors.New("1")
	}

	main, err := os.ReadFile(MainFile)

	if err != nil {
		fmt.Println("\nI can't find the main file")
		return errors.New("1")
	}

	mod, err := os.ReadFile("./go.mod")

	if err != nil {
		fmt.Println("\nI can't find the go.mod file")
		return errors.New("1")
	}

	var gomod string
	for x, i := range strings.Split(string(mod), "\n") {
		if x >= 1 {
			break
		}
		gomod = i[7:] // module name
	}

	var result string
	for _, text := range strings.Split(string(main), "\n") {
		result += text + "\n"
	}
	var module = `        "gomod/folder"`
	result = strings.ReplaceAll(result, "import (", `import (`+"\n"+strings.ReplaceAll(module, "gomod/folder", gomod+"/"+folder))

	ioutil.WriteFile(MainFile, []byte(result), os.ModePerm)
	return nil
}

// Generates random func
func randomFunc(fold string) []byte {

	var generated string
	var varname string

	generated += "\n\n"
	for i := 0; i < int(Random(1, 3)); i++ {
		generated += `//rand` + "\n" // random comment
		generated = strings.ReplaceAll(generated, "rand", RandomComment(10, 25))
	}

	funcname := UniqueFunc()
	generated += "func " + funcname + "() {\n"

	funcs = append(funcs, fold+"."+funcname)
	for i := 0; i < int(Random(2, 5)); i++ {
		varname = UniqueVar()
		generated += "var " + varname + randomVar() + "\n"
		generated += varname + " = " + varname + "\n"
	}

	if RandomArray(boolvalue) == "true" { // true or false

		generated += "\ngo func() {\n"
		generated += "\n\nfor i := 0; i < " + strconv.Itoa(int(Random(10, 10000))) + "; i++ {\n"

		for i := 0; i <= int(Random(1, 3)); i++ {
			generated += "\n\n" + `println("` + RandomString(5, 15) + `")`

		}
		generated += "\nprintln(i)\n"

		generated += "\n}\n"
		generated += "\n}()\n"

	}

	if RandomArray(boolvalue) == "false" {

		for i := 0; i <= int(Random(2, 5)); i++ {
			generated += "\nif " + strconv.Itoa(int(Random(1, 1000))) + ` == ` + strconv.Itoa(int(Random(1, 1000))) + " {\n"

			for x := 0; x <= int(Random(2, 5)); x++ {
				generated += "\n\n" + `println("` + RandomString(5, 15) + `")`
			}
			generated += "\n}\n"

		}

	}

	generated += "\n}"

	return []byte(generated)

}

func randomVar() string {

	temptype := RandomArray(vartypes)
	var data string

	if temptype == "byte" {
		data = `[]byte("rand")`
		data = strings.ReplaceAll(data, "rand", RandomString(5, 15))
	}
	if temptype == "int" {
		data = `rand`
		data = strings.ReplaceAll(data, "rand", strconv.Itoa(int(Random(100, 999999))))
	}
	if temptype == "string" {
		data = `"rand"`
		data = strings.ReplaceAll(data, "rand", RandomString(5, 15))
	}
	if temptype == "bool" {
		data = boolvalue[random.Intn(len(boolvalue))]
	}

	return " = " + data

}
