package simba

// YesterdayInfo 结构体
type YesterdayInfo struct {
	// 昨日点击量
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 昨日展现量
	Impression string `json:"impression,omitempty" xml:"impression,omitempty"`
}
