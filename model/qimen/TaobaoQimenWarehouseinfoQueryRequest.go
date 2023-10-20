package qimen

// TaobaoqimenwarehouseinfoqueryRequest 结构体
type TaobaoqimenwarehouseinfoqueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimenwarehouseinfoqueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
