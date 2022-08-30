package alihouse

// ProjectTradeItemOrderDto 结构体
type ProjectTradeItemOrderDto struct {
	// 交易工具图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 交易工具跳转地址
	JumpValue string `json:"jump_value,omitempty" xml:"jump_value,omitempty"`
	// 交易商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 类型
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}
