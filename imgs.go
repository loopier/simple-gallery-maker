// Create a very basic image gallery from images in subdirectories.
//
// Each subdirectory name will be used to create a section (a separate
// .html fiile) with the same name.
//
// For example, if there's a subdirectory named 'illustrations', a new
// 'illustrations.html' file will be created and the images in that directory
// will be displayed in alphabetical order.

package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"os"
)

// var imgsDir string = "illustration"
// var outputFile string = "illustration.html"
var headerFile string = "header.html"
var footerFile string = "footer.html"

func readFile( fileName string ) string {

    // fmt.Printf("\n\nReading a file in Go lang\n")
    // fileName := "test.txt"

    // The ioutil package contains inbuilt
    // methods like ReadFile that reads the
    // filename and returns the contents.
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    // fmt.Printf("\nFile Name: %s", fileName)
    // fmt.Printf("\nSize: %d bytes", len(data))
    // fmt.Printf("\nData: %s", data)
    // fmt.Printf("%s", data)
	return fmt.Sprintf("%s", data)
}

func createFile( fileName string, data string ) {

    // fmt package implements formatted
    // I/O and has functions like Printf
    // and Scanf
    fmt.Printf("Writing to a file in Go lang\n")

    // in case an error is thrown it is received
    // by the err variable and Fatalf method of
    // log prints the error message and stops
    // program execution
    file, err := os.Create(fileName)

    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

    // Defer is used for purposes of cleanup like
    // closing a running file after the file has
    // been written and main //function has
    // completed execution
    defer file.Close()

    // len variable captures the length
    // of the string written to the file.
    len, err := file.WriteString(data)

    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }

    // Name() method returns the name of the
    // file as presented to Create() method.
    fmt.Printf("\nFile Name: %s", file.Name())
    fmt.Printf("\nLength: %d bytes", len)
}

func createSection( name string ) {
    files, err := ioutil.ReadDir(name)
    if err != nil {
        log.Fatal(err)
    }

	var imgsStr string = readFile(headerFile)

    for _, file := range files {
        // fmt.Println(file.Name(), file.IsDir())
		imgsStr = fmt.Sprintf("%s<img src=\"%s/%s\" width=\"100%%\" />\n", imgsStr, name, file.Name())
    }

	imgsStr = fmt.Sprintf("%s%s", imgsStr, readFile(footerFile))

	// fmt.Println(readFile(headerFile))
	// fmt.Println(imgsStr)
	// fmt.Println(readFile(footerFile))

	var outputFile = fmt.Sprintf("%s.html", name)
	createFile( outputFile, imgsStr )

	fmt.Println(readFile(outputFile))
}

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			// fmt.Println(file.Name())
			if file.Name() != "fonts" {
				createSection( file.Name() )
			}
		}
	}
}
