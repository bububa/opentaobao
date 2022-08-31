package wdk

// BatchQueryRequest 结构体
type BatchQueryRequest struct {
	// 店铺号列表，order_from=4时必填，其他非必填
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 订单状态: PAID / PACKAGED / SUCCESS，非必填
	OrderStatus []string `json:"order_status,omitempty" xml:"order_status>string,omitempty"`
	// 下单终端: APP / POS，非必填
	OrderClient string `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 废弃字段
	SyncStatus string `json:"sync_status,omitempty" xml:"sync_status,omitempty"`
	// 起始时间，必填
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间，必填
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 渠道店id，order_from=4时非必填，其他必填
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店id，非必填
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 业务类型 aelophy/其它
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 页序号，从0开始，必填
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 单页大小，不超过200，必填
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 渠道来源，3：饿了么  4：盒马&amp;淘鲜达 18:大润发飞牛  19:欧尚外卖  20：商家自有渠道 ，必填
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}
