package jst

// SmsFileContentDto 结构体
type SmsFileContentDto struct {
	// 文件名称
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件后缀名
	FileSuffix string `json:"file_suffix,omitempty" xml:"file_suffix,omitempty"`
	// 文件Base64转码后的字符串
	FileContents string `json:"file_contents,omitempty" xml:"file_contents,omitempty"`
	// 文件大小
	FileSize int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
}
