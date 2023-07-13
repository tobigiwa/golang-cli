package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Buildig tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool, %v", err)
		os.Exit(1)
	}

	fmt.Println("Running test...")
	result := m.Run()
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)

}

func TestTodoCLI(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)
	task := "test task number 1"

	t.Run("AddNewTaskFromArguements", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		fmt.Println(cmd)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task2 := "test task number 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf("1: %s\n2: %s\n", task, task2)
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead \n", expected, string(out))
		}
	})
}
