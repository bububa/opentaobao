package tmallgenie

// AlarmMusic 
type AlarmMusic struct {
    // 铃声类型
    Category   string `json:"category,omitempty" xml:"category,omitempty"`
    // 铃声id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 铃声名字
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 铃声url
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 铃声来源，默认为虾米
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
}
