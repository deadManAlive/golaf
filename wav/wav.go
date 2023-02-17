package wav

import (
	"errors"
	"fmt"
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
}

func ReadFile(filename string) (Wav, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Wav{}, errors.New("error opening file")
	}

	word := make([]byte, 2)
	dword := make([]byte, 4)

	wav := Wav{}

	// format naming based on http://soundfile.sapp.org/doc/WaveFormat/
	// read chunkid
	f.Read(dword)
	chunkId, _ := util.ReadFourBytesBE(dword)
	fmt.Printf("chunkid: %#x ('%s')\n", chunkId, dword)

	// read format
	f.Seek(8, 0)
	f.Read(dword)
	format, _ := util.ReadFourBytesBE(dword)
	fmt.Printf("format: %#x ('%s')\n", format, dword)

	// read subchunkid1
	f.Seek(12, 0)
	f.Read(dword)
	subChunk1Id, _ := util.ReadFourBytesBE(dword)
	fmt.Printf("subchunk1id: %#x ('%s')\n", subChunk1Id, dword)

	// read audioformat
	f.Seek(20, 0)
	f.Read(word)
	audioFormat, _ := util.ReadTwoBytes(word)
	fmt.Printf("audioformat: %d ", audioFormat)
	switch audioFormat {
	case 0x0000:
		fmt.Println("(Unknown)")
	case 0x0001:
		fmt.Println("(PCM)")
	case 0xFFFF:
		fmt.Println("(Experimental)")
	default:
		fmt.Println("(Compressed formats)")
	}
	wav.audioFormat = audioFormat

	// read numchannels
	f.Seek(22, 0)
	f.Read(word)
	numChannels, _ := util.ReadTwoBytes(word)
	fmt.Printf("numch: %d\n", numChannels)
	wav.numChannels = numChannels

	// read samplerate
	f.Seek(24, 0)
	f.Read(dword)
	sampleRate, _ := util.ReadFourBytes(dword)
	fmt.Printf("samplerate: %d\n", sampleRate)
	wav.sampleRate = sampleRate

	// read byterate
	f.Seek(28, 0)
	f.Read(dword)
	byteRate, _ := util.ReadFourBytes(dword)
	fmt.Printf("byterate: %d\n", byteRate)
	wav.byteRate = byteRate

	// read blockalign
	f.Seek(32, 0)
	f.Read(word)
	blockAlign, _ := util.ReadTwoBytes(word)
	fmt.Printf("blockalign: %d\n", blockAlign)
	wav.blockAlign = blockAlign

	// read bitsperspl
	f.Seek(34, 0)
	f.Read(word)
	bitsPerSpl, _ := util.ReadTwoBytes(word)
	fmt.Printf("bitspersample: %d", bitsPerSpl)
	wav.bitsPerSpl = bitsPerSpl

	f.Close()
	return wav, nil
}
