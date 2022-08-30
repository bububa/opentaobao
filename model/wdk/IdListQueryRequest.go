package wdk

// IdListQueryRequest 结构体
type IdListQueryRequest struct {
	// 中台订单号
	BizIdList []int64 `json:"biz_id_list,omitempty" xml:"biz_id_list>int64,omitempty"`
	// 淘系订单号
	TbBizIdList []int64 `json:"tb_biz_id_list,omitempty" xml:"tb_biz_id_list>int64,omitempty"`
	// 渠道单号
	OutOrderIdList []string `json:"out_order_id_list,omitempty" xml:"out_order_id_list>string,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}
