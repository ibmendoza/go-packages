package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Declare variable b which is a byte array of
	// the string `Hello, World!`
	var b = []byte("Hello, World!")

	b = bytes.TrimSuffix(b, []byte(", World!"))

	// No change will be made here as `Gopher` is
	// not inside the byte slice b
	b = bytes.TrimSuffix(b, []byte("Gopher!"))

	// Write b to os.Stdout
	os.Stdout.Write(b)

	// Write a blank line
	fmt.Println()

	// Now append the byte slice `, World!` into b, but first
	// use bytes.TrimSuffix to trim any `!` suffix from `World!`
	b = append(b, bytes.TrimSuffix([]byte(", Gopher!"), []byte("!"))...)

	// Write b to os.Stdout again
	os.Stdout.Write(b)
}
