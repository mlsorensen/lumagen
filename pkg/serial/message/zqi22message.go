package message

type ZQI22Message struct {
	SourceAspectRatio uint8
	SourceFrameRate uint8
	SourceVerticalResolution uint16
	HDR bool
}
