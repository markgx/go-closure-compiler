package closure

import (
	"os/exec"
)

func Compile(files *[]string, outputFilePath string, options map[string]string) error {
	optionsArray := []string{}

	if options != nil {
		for k, v := range options {
			optionsArray = append(optionsArray, k)

			if v != "" {
				optionsArray = append(optionsArray, v)
			}
		}
	}

	args := make([]string, len(*files)*2+2)

	for i, file := range *files {
		args[i*2] = "--js"
		args[i*2+1] = file
	}

	args[len(args)-2] = "--js_output_file"
	args[len(args)-1] = outputFilePath

	args = append(optionsArray, args...)

	cmd := exec.Command("closure-compiler", args...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
