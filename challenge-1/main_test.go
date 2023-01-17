package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	testValue := "Hello, Galaxy"

	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage(testValue, &wg)
	wg.Wait()

	if !strings.Contains(msg, testValue) {
		t.Errorf("Expected to find '%s' but instead found '%s'", testValue, msg)
	}
}

func Test_printMessage(t *testing.T) {
	testValue := "Hello, Galaxy"

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = testValue
	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, testValue) {
		t.Errorf("Expected to find '%s' but instead found '%s'", testValue, output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but it is not there")
	}

}
