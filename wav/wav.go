package wav

import (
	"errors"
	"fmt"
	"os"

	"github.com/deadManAlive/golaf/util"
	"golang.org/x/exp/constraints"
)

func pow[T constraints.Integer](x, y T) T {
	var res T = 1
	for i := 0; i < int(y); i++ {
		res *= x
	}
	return res
}

// only works on mono, integer sample
func ReadFile(filename string) (Wav, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Wav{}, errors.New("error opening file")
	}
	defer f.Close()

	info, _ := f.Stat()
	size := info.Size()
	if size > int64(pow(2, 30)) {
		fmt.Printf("Warning: file size is massive (%v GB), expect performance issues.\n", size/1000000000)
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
	f.Read(word)
	numChannels, _ := util.ReadTwoBytes(word)
	res.numChannels = numChannels

	// read samplerate
	f.Read(dword)
	sampleRate, _ := util.ReadFourBytes(dword)
	res.sampleRate = sampleRate

	// read byterate
	f.Read(dword)
	byteRate, _ := util.ReadFourBytes(dword)
	res.byteRate = byteRate

	// read blockalign
	f.Read(word)
	blockAlign, _ := util.ReadTwoBytes(word)
	res.blockAlign = blockAlign

	// read bitsperspl
	f.Read(word)
	bitsPerSpl, _ := util.ReadTwoBytes(word)
	res.bitsPerSpl = bitsPerSpl

	// read subchunk2size
	f.Seek(40, 0)
	f.Read(dword)
	subChunk2Size, _ := util.ReadFourBytes(dword)
	bufLength := subChunk2Size * 8 / uint32(bitsPerSpl)
	buffer := make([]int, 0, bufLength)
	reader := make([]byte, bitsPerSpl/8)

	var s int
	for {
		_, e := f.Read(reader)
		if e != nil {
			break
		}
		s = util.LiToInt(reader) // TODO: small width buffer read as unsigned data
		buffer = append(buffer, s)
	}

	res.pcmSample = buffer
	return res, nil
}
