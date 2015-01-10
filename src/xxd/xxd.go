package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

const (
	byteOffsetInit = 8
	charOffsetInt  = 39
	line_length    = 50
)

func main() {
	line_offset := 0

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [file]\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf := make([]byte, 16)
		n, err := r.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}

		// Line offset
		fmt.Printf("%06d: ", line_offset)
		line_offset += 10

		// Hex values
		for i := 0; i < n; i += 2 {
			fmt.Printf("%02X ", buf[i:i+2])
		}
		for i := n; i < len(buf); i += 2 {
			fmt.Printf("     ")
		}

		fmt.Printf(" ")

		// Character values
		b := buf[:n]
		for len(b) > 0 {
			r, size := utf8.DecodeRune(b)

			if strconv.IsPrint(r) {
				fmt.Printf(string(r))
			} else {
				fmt.Printf(".")
			}
			b = b[size:]
		}

		fmt.Printf("\n")
	}
}
