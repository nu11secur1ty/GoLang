//nu11secur1ty
package main

import (
    "log"
    "os/exec"
)

func main() {

    cmd := exec.Command("path_/to/your/batman_filr/batman.bat")

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}

