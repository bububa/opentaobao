package product

// ItemImg 结构体
type ItemImg struct {
	// 图片链接地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 商品图片的id，和商品相对应（主图id默认为0）
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 图片放在第几张（多图时可设置）
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}
