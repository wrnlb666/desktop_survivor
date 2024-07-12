package util


import (
    "os"
    "os/exec"
)


func TaskNew(args ...string) {
    cmd := exec.Command(os.Args[0], args...)
    go cmd.Run()
}
