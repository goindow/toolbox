package toolbox

import (
	"testing"
	//"os"
	//"os/exec"
	//"strings"
)

/*const (
	ENV_DUMP  = "GO_TEST_TOOLBOX_DUMP"
)

func Test_Dump(t *testing.T) {
	if os.Getenv(ENV_DUMP) == "1" {
		Dump(os.Getwd())
		os.Exit(1)
	}
	cmd := exec.Command(os.Args[0], "-test.run=Test_Dump")
	cmd.Env = append(os.Environ(), ENV_DUMP+"=1")
	got, _ := cmd.Output()
	want := "[string] /usr/local/var/go/src/github.com/goindow/toolbox\n[<nil>] <nil>"
	if strings.TrimSpace(string(got)) != want {
		fail(t, "toolbox.Dump() should print info("+want+")")
	}
}*/

func fail(t *testing.T, s string) {
	echo(t, s, 1)
}

func ok(t *testing.T, s string) {
	echo(t, s, 2)
}

func echo(t *testing.T, s string, level uint) {
	switch level {
	case 1:
		t.Error("[fail] " + s)
	case 2:
		t.Log("[ok] " + s)
	}
}
