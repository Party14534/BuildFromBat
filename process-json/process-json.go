package processjson

import (
  "os"
  "fmt"
  "encoding/json"
  "io"
)

type CompileInfo struct {
  Compiler string
  Flags []string
  IncludeDirectories []string
  LibraryDirectories []string
  Libraries []string
  Excludes []string
  System string
  Extension string
}

func ProcessJson() CompileInfo {
  var info CompileInfo
  jsonFile, err := os.Open("project.json")
  
  if(err != nil) {
    fmt.Println("Error while opening project.json: ", err)
    os.Exit(1)
  }

  defer jsonFile.Close()

  jsonBytes, err := io.ReadAll(jsonFile)

  if(err != nil) { 
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
