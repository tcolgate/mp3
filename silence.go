package mp3

import (
	"bytes"

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
