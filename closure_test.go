package closure

import (
	"io/ioutil"
	"testing"
)

func compareOutput(t *testing.T, outputFile string, expectedString string) {
	// file outputFile should exist
	var d []byte
	d, err := ioutil.ReadFile(outputFile)
	s := string(d)
	if err != nil {
		t.Error(err)
	}

	if s != expectedString {
		t.Errorf("Minimized output does not match expected\nEXPECTED:\n%v\nACTUAL:\n%v", s, expectedString)
	}
}

func Test_Compile(t *testing.T) {
	// compiler single file
	Compile(&[]string{"fixtures/a.js"}, "fixtures/a.min.js", nil)
	expectedString := "function add(a,b){return sum=a+b};\n"
	compareOutput(t, "fixtures/a.min.js", expectedString)

	// compile two files
	Compile(&[]string{"fixtures/a.js", "fixtures/b.js"}, "fixtures/ab.min.js", nil)
	expectedString = "function add(a,b){return sum=a+b};function multiply(a,b,c){return a*b*c};\n"
	compareOutput(t, "fixtures/ab.min.js", expectedString)

	// set options
	options := make(map[string]string)
	options["--compilation_level"] = "WHITESPACE_ONLY"

	Compile(&[]string{"fixtures/a.js"}, "fixtures/a.min-whitespaceonly.js", options)
	expectedString = "function add(parameter1,parameter2){sum=parameter1+parameter2;return sum};\n"
	compareOutput(t, "fixtures/a.min-whitespaceonly.js", expectedString)
}
