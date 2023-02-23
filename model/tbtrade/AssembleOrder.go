package tbtrade

// AssembleOrder 结构体
type AssembleOrder struct {
	// 合单订单列表，一个列表最多200
	OrderList []OrderGroup `json:"order_list,omitempty" xml:"order_list>order_group,omitempty"`
	// 组合id，服务商内部的合单操作id，取消合单会根据group_id进行删除操作。
	GroupId string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}
