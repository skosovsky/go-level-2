package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	err := createFile()
	if err != nil {
		fmt.Println(err)
	}

	err = calculate()
	if err != nil {
		fmt.Println(err)
	}

}

func calculate() error {
	var userNum int

	fmt.Println("Введите 0 для ошибки деления на 0, или букву для других ошибок")
	_, err := fmt.Scan(&userNum)
	if err != nil {
		return MyError{time: fmt.Sprint(time.Now()), area: "calculate", innerError: err.Error()}
	}

	_, err = fmt.Println(10 / userNum)
	if err != nil {
		return MyError{time: fmt.Sprint(time.Now()), area: "calculate", innerError: err.Error()}
	}

	return nil
}

func createFile() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return MyError{time: fmt.Sprint(time.Now()), area: "create file", innerError: err.Error()}
	}
	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			return MyError{time: fmt.Sprint(time.Now()), area: "close file", innerError: err.Error()}
		}
		return nil
	}(f)
	return nil
}

type MyError struct {
	time       string
	area       string
	innerError string
}

func (err MyError) Error() string {
	return fmt.Sprintf("failed to %s: %s reason: %s", err.area, err.time, err.innerError)
}
