package examples

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestExampleThree(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ExampleThree()

	_ = w.Close()

	results, _ := io.ReadAll(r)
	output := string(results)

	os.Stdout = stdout

	if !strings.Contains(output, "34320.00") {
		t.Error("Wrong Balance Returned")
	}
}
