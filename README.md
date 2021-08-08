# fsearch
A cli for finding a file fast in a directory or a subdirectory of it in a parallel way.

The program will search all underlying files for the given file and show the paths to all the corresponding files.

## Installation
You can download the source code and build it with.
```
go build
```
Or download a compiled executable from the releases tab.

If you run
```
go install
```
In the installation directory you can skip step 2 in the usage.

## Usage
1. Open cmd
2. run 
```
cd <path to your installation folder>
```
3. run 
```
fsearch <path> <filename>
```
Path refers to the top folder in which you want to start searching.

Filename is the filename you're looking for.

4. Wait for the program to finish
5. Use the results

## Extras
All file searching functionality is kept in the fslib folder/package.
Because of this you can easily use the file searching functionality in any projects of your own.
