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
var Error error

/*
	Функция проверки существования пути
*/
//export check_Path
func check_Path(str *C.char) C.int {
	status, Error := os.Stat(C.GoString(str))
	if Error != nil {
		//Формируем сообщение о ошибке
		return 2 // Флаг ошибки
	}
	if status != nil {
		return 1
	} else {
		return 0
	}
}

/*
	Функция созлания файла
*/
//export createFile
func createFile(path *C.char) C.int {
	_, Error = os.Create(C.GoString(path))
	if Error != nil {
		return 0
	}
	return 1
}

/*
	Функция создания файла
*/
//export makeDir
func makeDir(path *C.char) C.int {
	Error = os.MkdirAll(C.GoString(path), 0777)
	if Error != nil {
		return 0
	}
	return 1
}

/*
	Функция удаления Файла или пустого Каталога
*/
//export removePath
func removePath(str *C.char) C.int {
	Error = os.Remove(C.GoString(str))
	if Error != nil {
		return 0
	}
	return 1
}

/*
	Функция удаления всего указаного пути
*/
//export removeAllPath
func removeAllPath(str *C.char) C.int {
	Error = os.RemoveAll(C.GoString(str))
	if Error != nil {
		return 0
	}
	return 1
}

/*
	Функция получения текущего пути
*/
//export getWd
func getWd(size *C.int) *C.char {
	dir, Error := os.Getwd()
	if Error != nil {
		return C.CString("Error")
	}
	*size = C.int(len(dir))
	return C.CString(dir)
}

/*
	Функция изменения имени текущего каталога
*/
//export changeDir
func changeDir(str *C.char) C.int {
	Error = os.Chdir(C.GoString(str))
	if Error != nil {
		return 0
	}
	return 1
}

/*
	Функция запуска команды на исполнения в текущем каталоге
*/
//export addCommand
func addCommand(path *C.char) *C.char {
	out, Error := exec.Command(C.GoString(path)).Output()
	if Error != nil {
		return C.CString(Error.Error())
	}
	return C.CString(string(out))
}

/*
	Функция запуска команды с одним аргументом на исполнения в текущем каталоге
*/
//export addCommands
func addCommands(command *C.char, args1 *C.char) *C.char {
	out, Error := commands(Dir, C.GoString(command), C.GoString(args1))
	if Error != nil {
		return C.CString(Error.Error())
	}
	return C.CString(string(out))
}

/*
	Функция смены текущего каталога для функций запуска команд
*/
//export setDir
func setDir(dir *C.char) {
	Dir = C.GoString(dir)
}

/*
Функция мониторина последней ошибки
*/
//export Check_Error
func Check_Error() *C.char {
	return C.CString(Error.Error())
}

/*
Выполнение команд
*/
func commands(dir string, comm string, args ...string) ([]byte, error) {
	cmd := exec.Command(comm, args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	return out, err
}

func main() {
	Dir = "./"
}
