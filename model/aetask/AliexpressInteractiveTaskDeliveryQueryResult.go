package aetask

// AliexpressInteractiveTaskDeliveryQueryResult 结构体
type AliexpressInteractiveTaskDeliveryQueryResult struct {
	// 物料集合
	Materials []Materials `json:"materials,omitempty" xml:"materials>materials,omitempty"`
	// 响应时间戳
	DateBegin string `json:"date_begin,omitempty" xml:"date_begin,omitempty"`
	// 风险等级
	SecurityLevel string `json:"security_level,omitempty" xml:"security_level,omitempty"`
	// 展示排序
	OrderDelivery int64 `json:"order_delivery,omitempty" xml:"order_delivery,omitempty"`
	// 分组id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 结束时间
	DateEnd int64 `json:"date_end,omitempty" xml:"date_end,omitempty"`
	// 预热展示标识
	PreDisplay bool `json:"pre_display,omitempty" xml:"pre_display,omitempty"`
}
