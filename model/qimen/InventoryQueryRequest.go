package qimen

// InventoryQueryRequest 
type InventoryQueryRequest struct {

    // 查询准则
    CriteriaList   []Criteria `json:"criteriaList,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

}
