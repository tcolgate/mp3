package mp3

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func BenchmarkDecode(t *testing.B) {
	// I suspect the compiler is actually just eliding
	// this
	r := bytes.NewReader(SilentBytes)
	d := NewDecoder(r)

	f := Frame{}
	d.Decode(&f)
}

func ExampleDecode() {
	r, err := os.Open("file.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := NewDecoder(r)
	var f Frame
	for {

		if err := d.Decode(&f); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(&f)
	}
}
