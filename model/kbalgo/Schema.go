package kbalgo

// Schema 结构体
type Schema struct {
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 页面类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// lbs信息
	Lbs string `json:"lbs,omitempty" xml:"lbs,omitempty"`
}
