package logistic

// PredictDeliveryTimeParam 结构体
type PredictDeliveryTimeParam struct {
	// 发件人地址
	TransportLocation *LocationParam `json:"transport_location,omitempty" xml:"transport_location,omitempty"`
	// 定位来源
	PositionSource string `json:"position_source,omitempty" xml:"position_source,omitempty"`
	// 餐品id
	IndexId string `json:"index_id,omitempty" xml:"index_id,omitempty"`
	// 收件人地址
	ReceiverLocation *LocationParam `json:"receiver_location,omitempty" xml:"receiver_location,omitempty"`
	// 配送费
	DeliveryTotalPrice int64 `json:"delivery_total_price,omitempty" xml:"delivery_total_price,omitempty"`
	// 餐品详情
	ItemInfos []ItemInfoParam `json:"item_infos,omitempty" xml:"item_infos>item_info_param,omitempty"`
	// appid
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 商户code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 门店code
	ChainStoreCode string `json:"chain_store_code,omitempty" xml:"chain_store_code,omitempty"`
	// 总数量
	OrderAmount int64 `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// 用户收件地址
	CustomerAddress string `json:"customer_address,omitempty" xml:"customer_address,omitempty"`
	// 预计送达时间
	RequireReceiveTime int64 `json:"require_receive_time,omitempty" xml:"require_receive_time,omitempty"`
}
