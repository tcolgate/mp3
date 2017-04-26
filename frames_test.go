package mp3

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkDecode(t *testing.B) {
	skipped := 0
	t.ReportAllocs()
	r := MakeSilence()
	d := NewDecoder(r)
	f := Frame{}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		t.SetBytes(int64(len(f.buf)))
		d.Decode(&f, &skipped)
	}
}

func ExampleDecoder_Decode() {
	skipped := 0
	r, err := os.Open("file.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := NewDecoder(r)
	var f Frame
	for {

		if err := d.Decode(&f, &skipped); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(&f)
	}
}
