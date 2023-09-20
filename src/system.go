package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import (
	"C"
	"os"
	"os/exec"
)

var Dir string

// Импортированные функции указаны с заглавной буквы
//
//export check_Path
func check_Path(str *C.char) C.int {
	_, err := os.Stat(C.GoString(str))
	if err != nil {
		//Формируем сообщение о ошибке
		return 0
	}
	return 1
}

//export createFile
func createFile(path *C.char) C.int {
	_, err := os.Create(C.GoString(path))
	if err != nil {
		return 0
	}
	return 1
}

//export makeDir
func makeDir(path *C.char) C.int {
	err := os.MkdirAll(C.GoString(path), 075)
	if err != nil {
		return 0
	}
	return 1
}

//export removePath
func removePath(str *C.char) C.int {
	err := os.Remove(C.GoString(str))
	if err != nil {
		return 0
	}
	return 1
}

//export removeAllInDir
func removeAllInDir(str *C.char) C.int {
	err := os.RemoveAll(C.GoString(str))
	if err != nil {
		return 0
	}
	return 1
}

//export getWd
func getWd(size *C.int) *C.char {
	dir, err := os.Getwd()
	if err != nil {
		return C.CString("Error")
	}
	*size = C.int(len(dir))
	return C.CString(dir)
}

//export changeDir
func changeDir(str *C.char) C.int {
	err := os.Chdir(C.GoString(str))
	if err != nil {
		return 0
	}
	return 1
}

//export addCommand
func addCommand(path *C.char) *C.char {
	out, err := exec.Command(C.GoString(path)).Output()
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString(string(out))
}

//export addCommands
func addCommands(command *C.char, args1 *C.char) *C.char {
	out, err := commands(Dir, C.GoString(command), C.GoString(args1))
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString(string(out))
}

//export setDir
func setDir(dir *C.char) {
	Dir = C.GoString(dir)
}

func commands(dir string, comm string, args ...string) ([]byte, error) {
	cmd := exec.Command(comm, args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	return out, err
}

func main() {
	Dir = "./"
}
