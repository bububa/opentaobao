package baoxian

// ClaimAttachment 结构体
type ClaimAttachment struct {
	// 附件类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 文件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 文件路径
	Path string `json:"path,omitempty" xml:"path,omitempty"`
	// 文件描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 文件类型
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
}
