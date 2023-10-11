package util

// FileDo 结构体
type FileDo struct {
	// 图片url
	FullUrl string `json:"full_url,omitempty" xml:"full_url,omitempty"`
	// 0
	ObjectKey string `json:"object_key,omitempty" xml:"object_key,omitempty"`
	// 上传文件夹id
	DirId int64 `json:"dir_id,omitempty" xml:"dir_id,omitempty"`
	// 图片大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}
