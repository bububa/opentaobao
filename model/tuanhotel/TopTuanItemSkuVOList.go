package tuanhotel

// TopTuanItemSkuVOList 结构体
type TopTuanItemSkuVOList struct {
	// 宝贝标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 套餐原价
	OrigPrice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 套餐价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 套餐名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 套餐间夜
	NightCount int64 `json:"night_count,omitempty" xml:"night_count,omitempty"`
	// 套餐库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 套餐人数
	PeopleCount int64 `json:"people_count,omitempty" xml:"people_count,omitempty"`
	// 日历库存信息
	CalendarInfo *TopSkuCalendarInfo `json:"calendar_info,omitempty" xml:"calendar_info,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
