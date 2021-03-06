package nrpos

// OfflineFileDto 结构体
type OfflineFileDto struct {
	// 文件名称
	FileKey string `json:"file_key,omitempty" xml:"file_key,omitempty"`
	// 文件下载地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
