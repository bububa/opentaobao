package ott

// PicCornerDo 结构体
type PicCornerDo struct {
	// 角标类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 角标地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 角标文案
	Text string `json:"text,omitempty" xml:"text,omitempty"`
}
