package perfect

// PerfectItemDescribeInfoDto 结构体
type PerfectItemDescribeInfoDto struct {
	// 详情描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 商品图片
	ItemPictures []string `json:"item_pictures,omitempty" xml:"item_pictures>string,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}
