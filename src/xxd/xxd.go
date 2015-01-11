package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	byteOffsetInit = 8
	charOffsetInt  = 39
	line_length    = 50
)

func main() {
	log := bufio.NewWriter(os.Stdout)
	defer log.Flush()

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
	buf := make([]byte, 16)
	for {
		n, err := io.ReadFull(r, buf)
		if n == 0 || err == io.EOF {
			break
		}

		// Line offset
		log.WriteString(fmt.Sprintf("%06x0: ", line_offset))
		line_offset++

		// Hex values
		for i := 0; i < n; i++ {
			log.WriteString(fmt.Sprintf("%02x", buf[i]))

			if i%2 == 1 {
				log.WriteByte(' ')
			}
		}
		if n < len(buf) {
			for i := n; i < len(buf); i++ {
				log.WriteString("  ")
				if i%2 == 1 {
					log.WriteByte(' ')
				}
			}
		}

		log.WriteString(" ")

		// Character values
		for i := 0; i < n; i++ {
			r := buf[i]
			if int(r) > 0x1f && int(r) < 0x7f {
				log.WriteString(fmt.Sprintf("%v", string(r)))
			} else {
				log.WriteByte('.')
			}
		}

		log.WriteByte('\n')
	}
}
