package wav

import (
	"errors"
	"os"

	"github.com/deadManAlive/golaf/util"
)

type Wav struct {
	audioFormat uint16
	numChannels uint16
	sampleRate  uint32
	byteRate    uint32
	blockAlign  uint16
	bitsPerSpl  uint16
	pcmSample   []float32
}

func ReadFile(filename string) (Wav, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Wav{}, errors.New("error opening file")
	}

	word := make([]byte, 2)
	dword := make([]byte, 4)

	res := Wav{}

	// format naming based on http://soundfile.sapp.org/doc/WaveFormat/
	// read chunkid
	f.Read(dword)
	chunkId, _ := util.ReadFourBytesBE(dword)

	// read format
	f.Seek(8, 0)
	f.Read(dword)
	format, _ := util.ReadFourBytesBE(dword)

	// read subchunkid1
	f.Seek(12, 0)
	f.Read(dword)
	subChunk1Id, _ := util.ReadFourBytesBE(dword)

	if chunkId != 0x52494646 || format != 0x57415645 || subChunk1Id != 0x666d7420 {
		return Wav{}, errors.New("unsupported or corrupted format")
	}

	// read audioformat
	f.Seek(20, 0)
	f.Read(word)
	audioFormat, _ := util.ReadTwoBytes(word)

	if audioFormat != 1 {
		return Wav{}, errors.New("non-LPCM format is unsupported")
	}

	res.audioFormat = audioFormat

	// read numchannels
	f.Seek(22, 0)
	f.Read(word)
	numChannels, _ := util.ReadTwoBytes(word)
	res.numChannels = numChannels

	// read samplerate
	f.Seek(24, 0)
	f.Read(dword)
	sampleRate, _ := util.ReadFourBytes(dword)
	res.sampleRate = sampleRate

	// read byterate
	f.Seek(28, 0)
	f.Read(dword)
	byteRate, _ := util.ReadFourBytes(dword)
	res.byteRate = byteRate

	// read blockalign
	f.Seek(32, 0)
	f.Read(word)
	blockAlign, _ := util.ReadTwoBytes(word)
	res.blockAlign = blockAlign

	// read bitsperspl
	f.Seek(34, 0)
	f.Read(word)
	bitsPerSpl, _ := util.ReadTwoBytes(word)
	res.bitsPerSpl = bitsPerSpl

	f.Close()
	res.pcmSample = []float32{1.0}
	return res, nil
}
