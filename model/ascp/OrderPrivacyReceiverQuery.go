package ascp

// OrderPrivacyReceiverQuery 结构体
type OrderPrivacyReceiverQuery struct {
	// 收件人
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 淘系交易单号
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 翱象强控出库单号 主场景：翱象下发的WMS出库单号  与交易单号必须二选一
	DeliveryOrderCode string `json:"delivery_order_code,omitempty" xml:"delivery_order_code,omitempty"`
	// 货主
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 使用场景。3001，即时配发货
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 店铺名称
	ShopNick string `json:"shop_nick,omitempty" xml:"shop_nick,omitempty"`
	// 业务请求时间，Unix timestamp ，毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
