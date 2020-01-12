package main

import (
	"fmt"
	"testing"
)

func TestGetVin(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(getVIN())
	}

}
