package legalsuit

// Attachmentlist 结构体
type Attachmentlist struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
