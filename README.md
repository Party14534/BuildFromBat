# Build From Bat

## Description
BuildFromBat simplifies the creation of build scripts for C++ projects, streamlining the development process.

BuildFromBat manages the discovery and manipulation of directories and files required for the build process which enhances the efficiency of the project setup and execution, contributing to a more streamlined workflow for C++ development.

## How to install
1. Ensure the latest version of Golang is installed
    - Golang can be installed from Go's official [website](https://go.dev/doc/install)
2. Clone the project repo 
3. Open the terminal in BuildFromBat's directory
4. Build the project using the go build command
```console
go build
```
6. Install BuildFromBat using go install
```console
go install
```
7. The program is now successfully installed, the BuildFromBat command should now be able to be used anywhere on your system!
    - If the program doesn't run from the BuildFromBat command ensure that the "go/bin" folder is in your systems PATH

## How to use
- **Run from Project's Root Directory:** Execute this tool from the highest-level directory of your project, the root directory.
- **Prepare `project.json`:** Ensure there's a `project.json` file within the directory. This file should detail essential information like compiler flags, required libraries, and other configurations necessary for your project's successful build.
  - *Note*: If 'project.json' isn't found in the execution directory, the tool will search for the global `project.json` in the '~/.config/BuildFromBat' directory.
- **Automated Build Setup:** BuildFromBat simplifies the process by creating a 'build' folder in the current directory if it doesn't exist already. It generates a .bat/.sh file within this folder. The resulting .bat/.sh file compiles the project into an executable with the desired name.

## Arguments
The program requires no arguments and can be run by simply calling its name

```console
BuildFromBat
```

If you wish to create a sample `project.json` in your current folder you can run:

```console
BuildFromBat -t
```

## `project.json` format
BuildFromBat uses the json provided by the user to fill out the info necessary to compile the project 

Every key except Compiler and Extension accept multiple occurrences in the `project.txt` file:

- **compiler:** Specifies the compiler invoked by the build process.
- **flags:** Indicates additional compiler flags required for the build.
- **includeDirectories:** Denotes the project's include directories.
- **libraryDirectories:** Specifies directories containing project libraries.
- **libraries:** Identifies libraries utilized by the compiler during the build.
- **excludes:** Identifies directories or files that should not be added to the build file. The exclude strings are interpreted as regex.
- **extension:** If this value is not set the program will choose the correct extension based on the user's operating system, however not all operating systems are currently supported.
- **name:** If this value is not set the program will set the name of the program to 'app'.

There is an example `project.json` file in this repo that you can build off of.

## Compiler Compatibility
Currently BuildFromBat only officially supports the `g++` compiler, however future updates will add functionality with other compilers.

## TODO
I plan to continuously improve and expand this project to improve my skills as a programmer. Planned updates include:

- ~**Regex support:** A future update will make the 'excludes' target files via regex to allow for more freedom in compiling projects.~

- ~**Change naming system:** Changed the way names were set to remove need for command line argument~

- ~**Create sample JSON:** I plan to add a feature into the program that, when ran, will create an empty project.json for the user to modify.~

- **Custom Include Paths:** Upcoming updates will introduce the capability to include paths that BuildFromBat will also scan to add files to the build script.
    - **Intended Use:** This feature aims to facilitate the inclusion of directories within your primary project. It will assist in scenarios where separate executable compilations within the project require files from sibling directories for compilation.

- **Extended Compiler Support:** Future updates will introduce support for additional compilers like `cl`, `gcc`, `clang`, and various other popular compilers, enhancing the tool's compatibility and versatility.

- **Expansion to Other Languages:** Long-term objectives involve extending support beyond C++ to encompass other compiled programming languages such as `C` and `Java`.
  
