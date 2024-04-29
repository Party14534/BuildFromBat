package BuildFromBat

import (
	createbat "BuildFromBat/create-bat"
	"BuildFromBat/filesystem"
	processjson "BuildFromBat/process-json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCompilerChanges(t *testing.T) {
    t.Run("Ensure bat file compiles with g++", func(t *testing.T) {
        project_info := processjson.CompileInfo {
            Compiler: "g++",
            Flags: []string{"Wall"},
            Libraries: []string{"Gl"},
            LibraryDirectories: []string{"path/to/lib"},
            IncludeDirectories: []string{"path/to/include"},
            Excludes: []string{"hidden"},
            Extension: ".sh",
        }
        
        parent := filesystem.NewDirectory(".", &project_info)

        createbat.WriteBat(&parent, &project_info, "test_bat")

        got_slice, err := os.ReadFile("build/build.sh")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        got := string(got_slice)

        base_filepath, err := filepath.Abs(".")
        if err != nil {
            t.Errorf("Unable to get absolute filepath")
            return
        }

        file1 := base_filepath + "/test_dir/file1.cpp"
        file2 := base_filepath + "/test_dir/file2.hpp"

        expected := `g++ -Wall \
-I"path/to/include" \
-L"path/to/lib" \
` + file2 + " \\\n" + file1 + " \\" + `
-o test_bat -lGl`

        if strings.TrimSpace(got) != strings.TrimSpace(expected) {
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

        got_slice, err := os.ReadFile("build/build.sh")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        got := string(got_slice)

        base_filepath, err := filepath.Abs(".")
        if err != nil {
            t.Errorf("Unable to get absolute filepath")
            return
        }

        file1 := base_filepath + "/test_dir/file1.cpp"
        file2 := base_filepath + "/test_dir/file2.hpp"

        expected := `gcc \
` + file2 + " \\\n" + file1 + " \\" + `
-o test_bat`

        if strings.TrimSpace(got) != strings.TrimSpace(expected) {
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

        got_slice, err := os.ReadFile("build/build.sh")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        got := string(got_slice)

        base_filepath, err := filepath.Abs(".")
        if err != nil {
            t.Errorf("Unable to get absolute filepath")
            return
        }

        file1 := base_filepath + "/test_dir/file1.cpp"
        file2 := base_filepath + "/test_dir/file2.hpp"

        expected := `g++ \
` + file2 + " \\\n" + file1 + ` \
-o test_bat`

        if strings.TrimSpace(got) != strings.TrimSpace(expected) {
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

        got_slice, err := os.ReadFile("build/build.bat")
        if err != nil {
            t.Errorf("Bat file was not created")
            return
        }

        got := string(got_slice)

        base_filepath, err := filepath.Abs(".")
        if err != nil {
            t.Errorf("Unable to get absolute filepath")
            return
        }

        file1 := base_filepath + "/test_dir/file1.cpp"
        file2 := base_filepath + "/test_dir/file2.hpp"

        expected := `g++ ^
` + file2 + " ^\n" + file1 + ` ^
-o test_bat`

        if strings.TrimSpace(got) != strings.TrimSpace(expected) {
            t.Errorf("Expected %v but got %v\n", expected, got)
        }
        
    })
}

