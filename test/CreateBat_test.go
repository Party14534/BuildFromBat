package BuildFromBat

import (
	createbat "BuildFromBat/create-bat"
	"BuildFromBat/filesystem"
	processjson "BuildFromBat/process-json"
	"os"
	"reflect"
	"testing"
)

func TestCompilerChanges(t *testing.T) {
    t.Run("Ensure bat file compiles with g++", func(t *testing.T) {
        project_info := processjson.CompileInfo {
            Compiler: "g++",
            Excludes: []string{"hidden"},
            Extension: ".sh",
        }
        
        parent := filesystem.NewDirectory(".", &project_info)

        createbat.WriteBat(&parent, &project_info, "test_bat")

        got, err := os.ReadFile("build/build.sh")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        expected_string := `g++
file1.cpp
file2.hpp
-o test_bat`
        expected := []byte(expected_string)

        if reflect.DeepEqual(got, expected) {
            t.Errorf("Expected %v but got %v\n", expected, got)
        }
        
    })

        t.Run("Ensure bat file compiles with gcc", func(t *testing.T) {
        project_info := processjson.CompileInfo {
            Compiler: "gcc",
            Excludes: []string{"hidden"},
            Extension: ".sh",
        }
        
        parent := filesystem.NewDirectory(".", &project_info)

        createbat.WriteBat(&parent, &project_info, "test_bat")

        got, err := os.ReadFile("build/build.sh")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        expected_string := `gcc
file1.cpp
file2.hpp
-o test_bat`
        expected := []byte(expected_string)

        if reflect.DeepEqual(got, expected) {
            t.Errorf("Expected %v but got %v\n", expected, got)
        }
        
    })
}

func TestExtensionChanges(t *testing.T) {
    t.Run("Ensure bat file works with .sh extension", func(t *testing.T) {
        project_info := processjson.CompileInfo {
            Compiler: "g++",
            Excludes: []string{"hidden"},
            Extension: ".sh",
        }
        
        parent := filesystem.NewDirectory(".", &project_info)

        createbat.WriteBat(&parent, &project_info, "test_bat")

        got, err := os.ReadFile("build/build.sh")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        expected_string := `g++
file1.cpp
file2.hpp
-o test_bat`
        expected := []byte(expected_string)

        if reflect.DeepEqual(got, expected) {
            t.Errorf("Expected %v but got %v\n", expected, got)
        }
        
    })

        t.Run("Ensure bat file works with .bat extension", func(t *testing.T) {
        project_info := processjson.CompileInfo {
            Compiler: "g++",
            Excludes: []string{"hidden"},
            Extension: ".bat",
        }
        
        parent := filesystem.NewDirectory(".", &project_info)

        createbat.WriteBat(&parent, &project_info, "test_bat")

        got, err := os.ReadFile("build/build.bat")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        expected_string := `g++ ^
file1.cpp ^
file2.hpp ^
-o test_bat`
        expected := []byte(expected_string)

        if reflect.DeepEqual(got, expected) {
            t.Errorf("Expected %v but got %v\n", expected, got)
        }
        
    })
}

