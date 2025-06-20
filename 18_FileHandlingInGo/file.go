package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// f, err := os.Open("sample.txt")

	// if err != nil {
	// 	panic(err) // You can use recover if panic needs to be handled
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name is : ", fileInfo.Name())

	// fmt.Println("Checking folder : ", fileInfo.IsDir())

	// fmt.Println("Size of file in bytes : ", fileInfo.Size())

	// fmt.Println("file permissions : ", fileInfo.Mode())

	// fmt.Println("File modified at : ", fileInfo.ModTime())

	// READ FILE -------------

	// f, err := os.Open("sample.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// // Make sure when you open file also try to close the file
	// defer f.Close()

	// // buffer is an array of bytes
	// // buf := [10]byte{}
	// buf := make([]byte, 100)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data in file is : ", d, string(buf[i]))
	// }

	// OS.ReadFile -------------NOTE A VIABLE SOLUTION IF FILE SIZE IS LARGE---------------

	// data, err := os.ReadFile("sample.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	//---------------------------- Read FOlders
	// dir, err := os.Open(".")

	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(5) // This number mentions how many files should be shown. -1 to show all files. Use '..' instead of number it will go back to a folder back

	// if err != nil {
	// 	panic(err)
	// }

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	// ---------------- Create a File

	// s, err := os.Create("sample2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer s.Close()

	// s.WriteString("Hi from Golang. Writing into a file")
	// s.WriteString("Nice language ")

	// ---------********************-------- For large File Reading-----------------
	//  			/// A KIND OF STREAMS fashion

	sourceFile, err := os.Open("sample.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	// bufio : This is a streaming kind of package
	reader := bufio.NewReader(sourceFile)
	//										// Default buff size is 4096 bytes
	writer := bufio.NewWriter(destFile)
	// Now do the pipe between reader and writer

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break // If EOF error is reached. So, we will get out of for infinite loop
		}

		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}

	}

	err = writer.Flush() // If any data is left we are flushing that data.

	if err != nil {
		panic(err)
	}

	log.Println("Written to new file successfully")

	// How to delete a file

	// err = os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
}
