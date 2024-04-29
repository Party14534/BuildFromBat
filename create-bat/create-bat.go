package createbat

import (
	"BuildFromBat/filesystem"
	"BuildFromBat/process-json"
	"fmt"
	"os"
)

func addAllFiles(contents *string, endLineChar string, dir *filesystem.Directory) {
    for _, file := range dir.Files {
        *contents += file + " " + endLineChar + "\n"
    }
    for _, subDirs := range dir.Directories {
        addAllFiles(contents, endLineChar, &subDirs)
    }
}

func WriteBat(parent *filesystem.Directory, 
info *processjson.CompileInfo, name string) {
    // If the build folder does not exist we create it
    _, err := os.Stat(parent.Directory.Name() + "/build")

    if os.IsNotExist(err) {
        err = os.Mkdir(parent.Directory.Name() + "/build", 0777)
        if err != nil {
            fmt.Println("Unable to create build folder")
            os.Exit(1)
        }
    }
    
    // Change the end line character based on the file extension
    endLineChar := "^"
    if info.Extension == ".sh" {
        endLineChar = "\\"
    }

    // String to hold the text before writing
    var contents string
    
    // Add the compiler
    contents += info.Compiler + " "
    
    // Add the flags from the json
    for _, flag := range info.Flags {
        contents += "-" + flag + " "
    }
    contents += endLineChar + "\n"
    
    // Add the include directores from the json
    for _, iDir := range info.IncludeDirectories {
        contents += "-I\"" + iDir + "\" "
    }

    // If there were any include directories start a new line
    if len(info.IncludeDirectories) > 0 { 
        contents += endLineChar + "\n" 
    }
    
    // Add the library directories
    for _, lDir := range info.LibraryDirectories {
        contents += "-L\"" + lDir + "\" "
    }

    // If there were any library directories start a new line
    if len(info.LibraryDirectories) > 0 { 
        contents += endLineChar + "\n" 
    }
    
    // Add all the file paths to the string
    addAllFiles(&contents, endLineChar, parent)

    // Add the -o flag and name of the program
    contents += "-o " + name + " "
    
    // Add the libraries to the end of the build script
    for _, lib := range info.Libraries {
        contents += "-l" + lib + " "
    }
    
    // Create or overwrite the build.bat file in the build folder
    err = os.WriteFile("build/build" + info.Extension, []byte(contents), 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
