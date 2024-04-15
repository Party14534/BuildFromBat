package processjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
)

type CompileInfo struct {
    Compiler string
    Flags []string
    IncludeDirectories []string
    LibraryDirectories []string
    Libraries []string
    Excludes []string
    Extension string
}

func ProcessJson() CompileInfo {
    var info CompileInfo
    
    // Open the project json or the global
    jsonFile, err := os.Open("project.json") 
    if err != nil {
        fmt.Println("Project level json not found, trying to use global project.json")
        usr, _ := user.Current()
        homeDir := usr.HomeDir
        
        jsonFile, err = os.Open(homeDir + "/.config/BuildFromBat/project.json")
        if err != nil {
            fmt.Println("Unable to open project.json file: ", err)
            os.Exit(1)
        }

        fmt.Println("Successfully used global project.json")
    }

    defer jsonFile.Close()

    jsonBytes, err := io.ReadAll(jsonFile)

    if err != nil { 
        fmt.Println("Error while converting to byte array") 
        os.Exit(1)
    }

    err = json.Unmarshal(jsonBytes, &info)

    if err != nil {
        fmt.Println("Error while decoding json: \n", err)
        os.Exit(1)
    } 

    return info
}
