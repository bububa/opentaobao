package ihome

// UploadPicMaterialDto 结构体
type UploadPicMaterialDto struct {
	// 图片宽
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 上传者ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 图片URL
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片高
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
