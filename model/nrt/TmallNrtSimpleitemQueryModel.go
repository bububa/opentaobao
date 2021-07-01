package nrt

// TmallNrtSimpleitemQueryModel 结构体
type TmallNrtSimpleitemQueryModel struct {
	// 0,1是正常状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商品主图
	ItemPic string `json:"item_pic,omitempty" xml:"item_pic,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品编码
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
