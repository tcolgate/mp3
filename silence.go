package mp3

import (
	"bytes"
	"log"
)

var (
	// SilentBytes is the raw data from a 128 Kb/s of lame encoded nothingness
	SilentBytes []byte
	// SilentFrame is the sound of Ripley screaming on the Nostromo, from the outside
	SilentFrame *Frame
)

//go:generate go-bindata -pkg mp3 -nomemcopy ./data
func init() {
	var err error
	SilentBytes, err = Asset("data/silent_1frame.mp3")
	if err != nil {
		log.Fatalf("Could not open silent_1frame.mp3 asset")
	}

	dec := NewDecoder(bytes.NewBuffer(SilentBytes))
	frame := Frame{}
	SilentFrame = &frame
	dec.Decode(&frame)
}
