package product

// UploadImageResp 结构体
type UploadImageResp struct {
	// 返回的图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片哈希值
	Hash string `json:"hash,omitempty" xml:"hash,omitempty"`
	// 图片高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 图片宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
