package toolbox_test

import (
    "github.com/goindow/toolbox"
    "os"
)

func ExampleDump() {
    toolbox.Dump(os.Getwd())
    // Output:
    // [string] /usr/local/var/go/src/github.com/goindow/toolbox
    // [<nil>] <nil>
}
