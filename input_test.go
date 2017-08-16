package input

import "testing"
import "strings"
import "bufio"

// TODO: better input tests
func TestPrompt(t *testing.T) {
	reader = bufio.NewReader(strings.NewReader("~/.non-existing-file\n~/.gitconfig\n"))

	out := Prompt("Test Question", RequiredValidator, FileValidator)
	t.Log("Out", out)
}
