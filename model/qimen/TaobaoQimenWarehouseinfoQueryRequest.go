package qimen

// TaobaoQimenWarehouseinfoQueryRequest 结构体
type TaobaoQimenWarehouseinfoQueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenWarehouseinfoQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
