package qimen

// InventoryQueryRequest 结构体
type InventoryQueryRequest struct {
	// 查询准则
	CriteriaList []Criteria `json:"criteriaList,omitempty" xml:"criteriaList>criteria,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenInventoryQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}
