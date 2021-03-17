package utils

import "os/exec"
import . "fmt"

func Command() {
    // app := "bin/bash"

	// arg0 := "export"
    // arg1 := "TEST=TRUE"

    // cmd := exec.Command(app, arg0, arg1)
    // cmd := exec.Command("sh", "ls")
    cmd := exec.Command("npm", "config", "set", "proxy", "http://username:password@host:port")
    stdout, err := cmd.Output()

    if err != nil {
        Println(err.Error())
        return
    }

    Print(string(stdout))
}