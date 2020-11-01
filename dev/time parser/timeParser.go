package main

import (
	"fmt"
	"time"
)

func HMS2UT(s string) int64 {
	//Nov  1 20:08:05
	t2, err := time.Parse("Jan  2 15:04:05", s)
	if err != nil {
		panic(err)
	}
	nix := t2.Unix()
	//fmt.Printf(">time:  %d\n", nix)
	timestamp := nix

	return timestamp
}
func UTC2UT(s string) int64 {

	var timestamp int64 = 0

	t2, err := time.Parse("2006-01-02T15:04:05.999999-07:00", s)
	if err != nil {
		panic(err)
	}
	nix := t2.Unix()
	//fmt.Printf(">time:  %d\n", nix)
	timestamp = nix

	return timestamp
}
func main() {
	stamp := "2020-11-01T17:17:52.840601+07:00"

	fmt.Printf("export:  %d\n", UTC2UT(stamp))
	fmt.Printf("export:  %d\n", HMS2UT("Nov  1 20:08:05"))
}
