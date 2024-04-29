package filesystem

import (
	"BuildFromBat/process-json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type Directory struct {
    Directory *os.File
    Directories []Directory
    Files []string
}

func (d Directory) DirectoryToString (str *string, offset int) {
    var tabs string = ""
    for i := 0; i < offset; i++ {
        tabs += "  "
    }

    for _, dir := range d.Directories {
        (*str) += tabs + "- " + filepath.Base(dir.Directory.Name()) + "\n"
        dir.DirectoryToString(str, offset + 1)
    }
    for _,file := range d.Files {
        (*str) += tabs + "* " + filepath.Base(file) + "\n"
    }
}

func (d Directory) PrintDirectory (offset int) {
    var tabs string = ""
    for i := 0; i < offset; i++ {
        tabs += "  "
    }

    for _, dir := range d.Directories {
        fmt.Println(tabs + "- " + filepath.Base(dir.Directory.Name()))
        dir.PrintDirectory(offset + 1)
    }
    for _,file := range d.Files {
        fmt.Println(tabs + "* " + filepath.Base(file))
    }
}

func NewDirectory(dirPath string, info *processjson.CompileInfo) Directory {
    var parent Directory

    directory, err := os.Open(dirPath)
    if(err != nil) { 
        fmt.Println(err)
        os.Exit(1)
    }
    parent.Directory = directory

    contents, err := parent.Directory.ReadDir(0)
    if(err != nil) {
        fmt.Println(err)
        os.Exit(1)
    }

    // Process contents and add to the parent directory
    for _, entry := range contents {
        entryInfo, err := entry.Info()
        if(err != nil) {
            fmt.Println(err)
            continue
        }

        // If the file or directory being added is in the exclude list, do not include them
        isExclude := false
        for _, exclude := range info.Excludes {
            r, err := regexp.Compile(`^` + exclude + `$`)
            if err != nil {
                fmt.Println("The following exclude is not a valid regex: ", exclude)
                os.Exit(1)
            }

            if r.MatchString(entry.Name()) {
                isExclude = true
                break
            } 
        }
        if isExclude { continue }

        absoultePath, _ := filepath.Abs(directory.Name() + "/" + entry.Name())
        ext := strings.Trim(path.Ext(absoultePath), " \n\r")

        if(entryInfo.IsDir()) {
            parent.Directories = append( 
                parent.Directories, 
                NewDirectory(absoultePath, info), 
            )
        } else if strings.Compare(ext, ".cpp") == 0  ||
        strings.Compare(ext, ".hpp") == 0 {
            parent.Files = append(parent.Files, absoultePath)
        }
    }

    return parent
}
