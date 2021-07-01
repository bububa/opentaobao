package maitix

// MoaOrderQueryParam 结构体
type MoaOrderQueryParam struct {
	// 大麦订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 查询是否排除取消的订单
	ExcludeUselessSubOrder bool `json:"exclude_useless_sub_order,omitempty" xml:"exclude_useless_sub_order,omitempty"`
}
