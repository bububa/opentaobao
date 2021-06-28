package qimen

// InventoryQueryRequest 
/* model for simplify = false
type InventoryQueryRequest struct {

    // 查询准则
    
    CriteriaList  struct {
        Criteria  []Criteria `json:"criteria,omitempty"`
    } `json:"criteriaList,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

}
*/

// InventoryQueryRequest 
type InventoryQueryRequest struct {

    // 查询准则
    CriteriaList   []Criteria `json:"criteriaList,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

}
