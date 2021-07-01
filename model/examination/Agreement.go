package examination

// Agreement 结构体
type Agreement struct {
	// 1代表图片2代表文件
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 二级制base64编码1
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}
