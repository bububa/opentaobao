package trade

// AdvertiseInfoQuery 结构体
type AdvertiseInfoQuery struct {
	// 用户id
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
