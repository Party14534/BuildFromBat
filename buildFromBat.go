package main

import (
	"fmt"
	"buildFromBat/create-bat"
	"buildFromBat/filesystem"
	"buildFromBat/process-json"
	"os"
	"runtime"
	"strings"
)

func main() {
    args := os.Args[1:]
    
    if len(args) != 1 {
        fmt.Println("Input the name of the output file as a command-line argument")
        os.Exit(1)
    }

    info := processjson.ProcessJson()
    
    // Set extension variable if not set by user
    userOs := runtime.GOOS
    if info.Extension == "" {
        if strings.Compare(userOs, "windows") == 0 {
            info.Extension = ".bat"
        } else if strings.Compare(userOs, "linux") == 0 {
            info.Extension = ".sh"
        } else { // Throw error if unsupported OS
            fmt.Println("Operating system is unsupported")
            os.Exit(1)
        }
    }

    parent := filesystem.NewDirectory(".", &info)
    parent.PrintDirectory(0)

    createbat.WriteBat(&parent, &info, args[0])
}
