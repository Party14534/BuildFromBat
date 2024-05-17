package main

import (
	"BuildFromBat/create-bat"
	createtemplate "BuildFromBat/create-template"
	"BuildFromBat/filesystem"
	"BuildFromBat/process-json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func main() {
    args := os.Args[1:]
    
    if len(args) == 1 {
        if reflect.DeepEqual(args[0], "-t") {
            fmt.Println("Creating process.json template")
            createtemplate.CreateTemplate()
        } else {
            fmt.Println("Invalid argument")
            os.Exit(1)
        }
        
        return
    }

    info, err := processjson.ProcessJson(true)
    if err != nil {
        os.Exit(1)
    }
    
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

    createbat.WriteBat(&parent, &info)
}
