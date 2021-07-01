package ott

// SourceInfo 结构体
type SourceInfo struct {
	// v1080
	V1080 string `json:"v1080,omitempty" xml:"v1080,omitempty"`
	// v720
	V720 string `json:"v720,omitempty" xml:"v720,omitempty"`
	// v480
	V480 string `json:"v480,omitempty" xml:"v480,omitempty"`
	// v320
	V320 string `json:"v320,omitempty" xml:"v320,omitempty"`
	// v240
	V240 string `json:"v240,omitempty" xml:"v240,omitempty"`
	// vBlueray4k
	VBlueray4k string `json:"v_blueray4k,omitempty" xml:"v_blueray4k,omitempty"`
	// hlsContentUrl
	HlsContentUrl string `json:"hls_content_url,omitempty" xml:"hls_content_url,omitempty"`
	// v2160tv
	V2160tv string `json:"v2160tv,omitempty" xml:"v2160tv,omitempty"`
	// v1080tv
	V1080tv string `json:"v1080tv,omitempty" xml:"v1080tv,omitempty"`
	// v720tv
	V720tv string `json:"v720tv,omitempty" xml:"v720tv,omitempty"`
	// v480tv
	V480tv string `json:"v480tv,omitempty" xml:"v480tv,omitempty"`
	// v320tv
	V320tv string `json:"v320tv,omitempty" xml:"v320tv,omitempty"`
	// v240tv
	V240tv string `json:"v240tv,omitempty" xml:"v240tv,omitempty"`
}
