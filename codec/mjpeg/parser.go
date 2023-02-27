package mjpeg

import "github.com/naresh-kavali/vdk/av"

type CodecData struct {
}

func (d CodecData) Type() av.CodecType {
	return av.MJPEG
}
