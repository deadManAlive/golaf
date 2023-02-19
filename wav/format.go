package wav

import "fmt"

type Wav struct {
	audioFormat uint16
	numChannels uint16
	sampleRate  uint32
	byteRate    uint32
	blockAlign  uint16
	bitsPerSpl  uint16
	pcmSample   []int
}

func (w Wav) String() string {
	res := fmt.Sprintf("audio format: %v\n", w.audioFormat)
	res += fmt.Sprintf("num channels: %v\n", w.numChannels)
	res += fmt.Sprintf("sample rate : %v\n", w.sampleRate)
	res += fmt.Sprintf("byte rate   : %v\n", w.byteRate)
	res += fmt.Sprintf("block align : %v\n", w.blockAlign)
	res += fmt.Sprintf("bit depth   : %v\n", w.bitsPerSpl)
	res += fmt.Sprintf("samples (IL): %v", w.pcmSample[:3])
	return res
}
