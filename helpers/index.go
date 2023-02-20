package helpers

type PathOptions struct {
	Src  string
	Dest string
}

type ProcessorOptions struct {
	MaximumDimensionPx int
	ImgQuality         int
}

type FileMap map[string][]byte

type String string
