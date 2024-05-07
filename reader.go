package main

import (
	"bufio"
	"bytes"
	"os"
)

func reader() ([]byte, error) {
	var bb bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Bytes()

		if len(text) < 0 {
			break
		}

		bb.Write(text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}
