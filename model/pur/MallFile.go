package pur

// MallFile 结构体
type MallFile struct {
	// 附件编码
	FileMd5 string `json:"file_md5,omitempty" xml:"file_md5,omitempty"`
	// 附件名称
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 附件URL
	FileUrl string `json:"file_url,omitempty" xml:"file_url,omitempty"`
}
