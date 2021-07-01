package alihealthmedical

// UploadImageResponse 结构体
type UploadImageResponse struct {
	// 文件key值
	ObjectKey string `json:"object_key,omitempty" xml:"object_key,omitempty"`
	// url
	FullUrl string `json:"full_url,omitempty" xml:"full_url,omitempty"`
}
