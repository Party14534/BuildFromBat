package BuildFromBat

import (
	"BuildFromBat/process-json"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestJsonProcessing(t *testing.T) {
    t.Run("No error", func(t * testing.T) { 
        contents := `{
            "compiler": "g++",
            "flags": [
                "Wall",
                "fsanitize=address"
            ],
            "includeDirectories": [
                "path/to/include"
            ],
            "libraryDirectories": [
                "path/to/library"
            ],
            "libraries": [
                "Gl",
                "sfml-window"
            ],
            "excludes": [
                "Test*"
            ],
            "extension": ".sh"
        }
        `

        error := os.WriteFile("project.json", []byte(contents), 0644)
        if error != nil {
            t.Errorf("Failed to write to project.json")
            return
        }


        got, err := processjson.ProcessJson() 
        if err != nil {
            t.Errorf("Error was thrown when processing json")
            return
        }

        expected := processjson.CompileInfo {
            Compiler: "g++",
            Flags: []string{"Wall", "fsanitize=address"},
            IncludeDirectories: []string{"path/to/include"},
            LibraryDirectories: []string{"path/to/library"},
            Libraries: []string{"Gl", "sfml-window"},
            Excludes: []string{"Test*"},
            Extension: ".sh",
        }

        if !reflect.DeepEqual(got, expected) {
            t.Errorf("Expected %v but got %v\n", expected, got)
        }
    })
    t.Run("Ensure error on invalid json", func(t * testing.T) { 
        
        contents := `{
            "compiler": "g++",
            "excludes": []
            "extension": ".sh"
        }`

         error := os.WriteFile("project.json", []byte(contents), 0644)
         if error != nil {
             t.Errorf("Failed to write to project.json")
             return
         }

        _, err := processjson.ProcessJson() 

        if err == nil {
            t.Errorf("Sent a bad json file and no error was thrown")
        }
    })
    t.Run("Invalid json fuzzing", func(t * testing.T) { 
        
        contents := "{"
        for i := 0; i < rand.Int() % 1000; i++ {
            contents += string('A' + rune(rand.Intn(60)))
        }

         error := os.WriteFile("project.json", []byte(contents), 0644)
         if error != nil {
             t.Errorf("Failed to write to project.json")
             return
         }

        _, err := processjson.ProcessJson() 

        if err == nil {
            t.Errorf("Sent a bad json file and no error was thrown")
        }
    })
}
