package ascp

// MediaResourceDto 结构体
type MediaResourceDto struct {
	// 图片/视频名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片/视频相对地址
	Path string `json:"path,omitempty" xml:"path,omitempty"`
}
