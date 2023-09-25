package main

import (
	"fmt"
	"testing"
)

func TestBar(t *testing.T) {
	res := Bar()
	if res != "good" {
		t.Errorf("expecting good, got %s", res)
	}
	fmt.Println(res)
}
