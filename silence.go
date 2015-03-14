package mp3

import (
	"bytes"
	"io"

	"github.com/tcolgate/mp3/internal/data"
)

var (
	// SilentFrame is the sound of Ripley screaming on the Nostromo, from the outside
	SilentFrame *Frame
	SilentBytes []byte
)

//go:generate go-bindata -pkg mp3 -nomemcopy ./data
func init() {
	SilentBytes = data.SilentBytes

	dec := NewDecoder(bytes.NewBuffer(SilentBytes))
	frame := Frame{}
	SilentFrame = &frame
	dec.Decode(&frame)
}

type silenceReader struct {
	int // Location into the silence frame
}

func (s *silenceReader) Close() error {
	return nil
}

func (s *silenceReader) Read(out []byte) (int, error) {
	for i := 0; i < len(out); i++ {
		out[i] = SilentBytes[s.int]
		s.int += 1
		if s.int >= len(SilentBytes) {
			s.int = 0
		}
	}

	return len(out), nil
}

func MakeSilence() io.ReadCloser {
	return &silenceReader{0}
}
