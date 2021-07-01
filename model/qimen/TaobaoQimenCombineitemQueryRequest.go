package qimen

// TaobaoQimenCombineitemQueryRequest 结构体
type TaobaoQimenCombineitemQueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenCombineitemQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
