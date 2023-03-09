package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Foo() {
	fmt.Println("Hi, This is hello package")
}

func RunModelUploadPyScript() {
	var PackagePath, _ = os.Getwd()
	var PythonPackage = "pymodule"
	var PythonScript = filepath.Join(PythonPackage, "main.py")
	var PythonInterpreter = filepath.Join(PythonPackage, ".venv/bin/python")
	// var PythonInterpreter = "pymodule/.venv/bin/python"

	fmt.Printf("PackagePath: %s\n", PackagePath)
	fmt.Printf("PythonPackage: %s\n", PythonPackage)
	fmt.Printf("PythonScript: %s\n", PythonScript)
	fmt.Printf("PythonInterpreter: %s\n", PythonInterpreter)
	cmd := exec.Command(PythonInterpreter, PythonScript)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Errorf("Failed to start model upload script: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Errorf("Failed to start model upload script: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Errorf("Failed to start model upload script: %v", err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)

	// Wait waits for the command to exit and
	// waits for any copying to stdin or copying from stdout or stderr to complete
	cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("%v\n",scanner.Text())
	}
}

func main() {
	Foo()
	RunModelUploadPyScript()
}
