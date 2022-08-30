package alitripmerchant

// EventParam 结构体
type EventParam struct {
	// 活动 ID
	OfferIdList []int64 `json:"offer_id_list,omitempty" xml:"offer_id_list>int64,omitempty"`
	// 订单 ID
	OrderIdList []string `json:"order_id_list,omitempty" xml:"order_id_list>string,omitempty"`
	// 酒店 ID
	HotelIdList []string `json:"hotel_id_list,omitempty" xml:"hotel_id_list>string,omitempty"`
	// 触发事件
	Event string `json:"event,omitempty" xml:"event,omitempty"`
	// 用户 token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
}
