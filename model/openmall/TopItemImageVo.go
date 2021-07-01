package openmall

// TopItemImageVo 结构体
type TopItemImageVo struct {
	// 图片放在第几张（多图时设置），propertyImages中不返回
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
	// 商品图片链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片ID，itemImages中不返回
	ImageId int64 `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 商品属性，itemImages中不返回
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
}
