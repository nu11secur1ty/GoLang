// @nu11secur1ty 2023
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
	
    f, err := os.Create("alien.bat")

    if err != nil {
        log.Fatal(err)
    }
	
    defer f.Close()

    val := `@echo off
notepad`

    data := []byte(val)

    _, err2 := f.Write(data)

    if err2 != nil {
        log.Fatal(err2)
    }
		fmt.Println("done..The file was created...")
}
