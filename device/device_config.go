package device

import (
	"github.com/vladimirvivien/go4vl/v4l2"
)

type config struct {
	ioType      v4l2.IOType
	pixFormat   v4l2.PixFormat
	bufSize     uint32
	fps         uint32
	bufType     uint32
	captureMode uint32
	inputIndex  int32
}

type Option func(*config)

func WithIOType(ioType v4l2.IOType) Option {
	return func(o *config) {
		o.ioType = ioType
	}
}

func WithPixFormat(pixFmt v4l2.PixFormat) Option {
	return func(o *config) {
		o.pixFormat = pixFmt
	}
}

func WithBufferSize(size uint32) Option {
	return func(o *config) {
		o.bufSize = size
	}
}

func WithFPS(fps uint32) Option {
	return func(o *config) {
		o.fps = fps
	}
}

func WithCaptureMode(captureMode uint32) Option {
	return func(o *config) {
		o.captureMode = captureMode
	}
}

func WithVideoCaptureEnabled() Option {
	return func(o *config) {
		o.bufType = v4l2.BufTypeVideoCapture
	}
}

func WithVideoOutputEnabled() Option {
	return func(o *config) {
		o.bufType = v4l2.BufTypeVideoOutput
	}
}

func WithVideoInputIndex(index int32) Option {
	return func(o *config) {
		o.inputIndex = index
	}
}
