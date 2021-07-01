package moscm

// ProductImgDto 结构体
type ProductImgDto struct {
	// 图片ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 排序
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
	// url地址（以http或https开头的绝对路径）
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
