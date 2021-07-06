package trade

// OrderDiscountBillQueryRequest 结构体
type OrderDiscountBillQueryRequest struct {
	// 查询结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 查询开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 经营店id，store_id和out_shop_code不能同时为空
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 外部门店编码
	OutShopCode string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
	// 传入上一次查询结果的next_id，第一次查询时传0
	NextId int64 `json:"next_id,omitempty" xml:"next_id,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 订单渠道 1代表"轻POS",2代表"淘鲜达",默认为 1
	OrderChannel int64 `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
}
