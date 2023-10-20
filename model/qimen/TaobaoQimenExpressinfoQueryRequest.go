package qimen

// TaobaoqimenexpressinfoqueryRequest 结构体
type TaobaoqimenexpressinfoqueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimenexpressinfoqueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
