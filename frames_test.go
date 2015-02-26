package mp3

import (
	"bytes"
	"log"
	"testing"
)

func BenchmarkDecode(t *testing.B) {
	r := bytes.NewReader(SilentBytes)
	d := NewDecoder(r)

	f := Frame{}
	d.Decode(&f)

	log.Println(&f)
}
