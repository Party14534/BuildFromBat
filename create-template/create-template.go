package createtemplate

import (
	"fmt"
	"os"
)

func CreateTemplate() {
    // Create a template project.json
    template := `{
    "compiler": "",

    "flags": [""],
  
    "includeDirectories": [""],
  
    "libraryDirectories": [""],
  
    "libraries": [""],
  
    "excludes": [""],
  
    "extension": ".sh",
  
    "name": "template"
}`
    
    err := os.WriteFile("project.json",[]byte(template), 0644);
    if err != nil {
        fmt.Println("Error creating template: ", err)
       os.Exit(1)
    }
}
