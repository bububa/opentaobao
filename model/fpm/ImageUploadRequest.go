package fpm

// ImageUploadRequest 结构体
type ImageUploadRequest struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件字节数组Base64字符串
	FileByteBase64Str string `json:"file_byte_base64_str,omitempty" xml:"file_byte_base64_str,omitempty"`
	// 文件大小
	FileSize int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
}
