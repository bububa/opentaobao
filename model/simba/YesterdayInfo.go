package simba

// YesterdayInfo 
type YesterdayInfo struct {
    // 昨日点击量
    Click   string `json:"click,omitempty" xml:"click,omitempty"`
    // 昨日展现量
    Impression   string `json:"impression,omitempty" xml:"impression,omitempty"`
}
