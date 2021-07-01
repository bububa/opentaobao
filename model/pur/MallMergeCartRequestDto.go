package pur

// MallMergeCartRequestDto 结构体
type MallMergeCartRequestDto struct {
	// 外部平台标识
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// 合并的商品行
	Items []MallItemDto `json:"items,omitempty" xml:"items>mall_item_dto,omitempty"`
	// 外部订单标识
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 用户加密id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
