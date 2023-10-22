package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(buff []byte) (int, error) {
	n, err := rr.r.Read(buff)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		c := buff[i]
		switch {
			case 'A' <= c && c <= 'Z':
			buff[i] = (c-'A'+13)%26 + 'A'
			case 'a' <= c && c <= 'z':
			buff[i] = (c-'a'+13)%26 + 'a'
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
