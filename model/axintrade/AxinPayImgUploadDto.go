package axintrade

// AxinPayImgUploadDto 结构体
type AxinPayImgUploadDto struct {
	// 图片类型
	ImageType string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
}
