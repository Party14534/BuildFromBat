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

func ProcessJson(use_global_json bool) (CompileInfo, error) {
    var info CompileInfo
    
    // Open the project json or the global
    jsonFile, err := os.Open("project.json") 
    if err != nil {
        fmt.Println("Project level json not found, trying to use global project.json")
        usr, _ := user.Current()
        homeDir := usr.HomeDir
        
        jsonFile, err = os.Open(homeDir + "/.config/BuildFromBat/project.json")
        if err != nil || !use_global_json {
            error := fmt.Errorf("Unable to open project.json file: %v\n", err)
            return info, error
        }

        fmt.Println("Successfully used global project.json")
    }

    defer jsonFile.Close()

    jsonBytes, err := io.ReadAll(jsonFile)

    if err != nil { 
        error := fmt.Errorf("Error while converting to byte array: %v\n", err) 
        return info, error
    }

    err = json.Unmarshal(jsonBytes, &info)

    if err != nil {
        error := fmt.Errorf("Error while decoding json: %v\n", err)
        return info, error
    } 

    return info, nil
}
