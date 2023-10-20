package qimen

// InventoryQueryRequest 结构体
type InventoryQueryRequest struct {
	// 查询准则
	CriteriaList []Criteria `json:"criteriaList,omitempty" xml:"criteriaList>criteria,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimeninventoryqueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
