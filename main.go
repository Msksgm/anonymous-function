package main

import (
	"fmt"
	"log"
)

func throwErr(flag bool) error {
	if flag {
		return fmt.Errorf("err is occured")
	} else {
		return nil
	}
}

func returnErr(flag bool) (err error) {
	defer func() {
		err = throwErr(flag)
	}()
	return nil
}

func main() {
	if err := returnErr(true); err != nil {
		log.Fatal(err)
	}
}
