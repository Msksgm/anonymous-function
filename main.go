package main

import (
	"fmt"
	"log"
)

func throwErr(flag bool) error {
	fmt.Println("throwErr")
	if flag {
		return fmt.Errorf("err is occured in throwErr")
	} else {
		return nil
	}
}

func returnErr(flag bool) (err error) {
	fmt.Println("returnErr")
	defer func() {
		err = throwErr(flag)
	}()
	return nil
}

func main() {
	fmt.Println("main function is start")
	if err := returnErr(true); err != nil {
		log.Fatal(err)
	}
	fmt.Println("main function is end")
}
