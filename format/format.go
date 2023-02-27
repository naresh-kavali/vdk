package format

import (
	"github.com/naresh-kavali/vdk/av/avutil"
	"github.com/naresh-kavali/vdk/format/aac"
	"github.com/naresh-kavali/vdk/format/flv"
	"github.com/naresh-kavali/vdk/format/mp4"
	"github.com/naresh-kavali/vdk/format/rtmp"
	"github.com/naresh-kavali/vdk/format/rtsp"
	"github.com/naresh-kavali/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
