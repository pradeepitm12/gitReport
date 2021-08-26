package generate

import (
	"github.com/pradeepitm12/gitReport/pkg/test"
	"testing"
)

func TestGitWork_invalid(t *testing.T) {
	gitwork := Command()
	out, err := test.ExecuteCommand(gitwork, "f")
	if err == nil {
		t.Errorf("No errors was defined. Output: %s", out)
	}
}
