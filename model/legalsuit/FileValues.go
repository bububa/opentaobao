package legalsuit

// FileValues 结构体
type FileValues struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 预览地址
	PreviewUrl string `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	// 下载地址
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// 文件类型
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	// 文件id
	FileId int64 `json:"file_id,omitempty" xml:"file_id,omitempty"`
}
