package flight

// UploadFileInfoDto 结构体
type UploadFileInfoDto struct {
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
