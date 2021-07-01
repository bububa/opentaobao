package openmall

// TopTradeResultVo 结构体
type TopTradeResultVo struct {
	// 淘宝交易ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 发货地址对应的areaid
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 发货地址
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 运费列表
	Posts []PostDo `json:"posts,omitempty" xml:"posts>post_do,omitempty"`
}
