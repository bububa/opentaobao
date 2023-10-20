package qimen

// TaobaoqimencombineitemqueryRequest 结构体
type TaobaoqimencombineitemqueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimencombineitemqueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
