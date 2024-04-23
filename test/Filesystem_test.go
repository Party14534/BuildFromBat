package BuildFromBat

import (
	"BuildFromBat/filesystem"
	processjson "BuildFromBat/process-json"
	"reflect"
	"testing"
)

func TestFilesystem(t * testing.T) {
    t.Run("Ensure filesystem adds all files and folders", func(t *testing.T) {
        project_info := processjson.CompileInfo { }
        
        got := new(string)
        filesystem.NewDirectory(".", &project_info).DirectoryToString(got, 0)
        expected := `- test_dir
  - hidden
    * hidden.hpp
    * hidden.cpp
  * file2.hpp
  * file1.cpp
  `

        if reflect.DeepEqual(*got, expected) {
            t.Errorf("Expected\n %v but got\n %v\n", expected, *got)
        }
    })

    t.Run("Ensure simple excludes work", func(t *testing.T) {
        project_info := processjson.CompileInfo { 
            Excludes: []string{"hidden"},
        }
        
        got := new(string)
        filesystem.NewDirectory(".", &project_info).DirectoryToString(got, 0)
        expected := `- test_dir
  * file2.hpp
  * file1.cpp
  `

        if reflect.DeepEqual(*got, expected) {
            t.Errorf("Expected\n %v but got\n %v\n", expected, *got)
        }
    })

    t.Run("Ensure simple regex excludes work", func(t *testing.T) {
        project_info := processjson.CompileInfo { 
            Excludes: []string{"*.hpp"},
        }
        
        got := new(string)
        filesystem.NewDirectory(".", &project_info).DirectoryToString(got, 0)
        expected := `- test_dir
  - hidden
    * hidden.cpp
  * file1.cpp
  `

        if reflect.DeepEqual(*got, expected) {
            t.Errorf("Expected\n %v but got\n %v\n", expected, *got)
        }
    })

}
