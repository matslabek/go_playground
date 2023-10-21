package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	// Fill the byte slice with 'A' characters.
	for i := range b {
		b[i] = 'A'
	}
	// Return the number of bytes read and a nil error since this reader never returns an error.
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}