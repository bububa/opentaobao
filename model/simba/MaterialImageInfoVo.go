package simba

// MaterialImageInfoVo 结构体
type MaterialImageInfoVo struct {
	// 图片链接
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 物料链接
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 物料名称
	MaterialName string `json:"material_name,omitempty" xml:"material_name,omitempty"`
	// 物料id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 主副图图片位置，主图0,副图从1开始
	ImagePosition int64 `json:"image_position,omitempty" xml:"image_position,omitempty"`
}
