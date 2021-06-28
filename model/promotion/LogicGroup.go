package promotion

// LogicGroup 
type LogicGroup struct {

    // 参与者分组序号
    
    Number   int64 `json:"number,omitempty" xml:"number,omitempty"`
    

    // 五道口参与者名称
    
    WdkGroupName   string `json:"wdk_group_name,omitempty" xml:"wdk_group_name,omitempty"`
    

    // 逻辑分组类型  COMMON(1, "普通分组"), EXCHANGE(2, "换购分组"), BUY_GIFT(3, "买赠分组"), EXCHANGE_TJ_OVERLAY(4, "特价换购叠加分组"),
    
    LogicGroupType   int64 `json:"logic_group_type,omitempty" xml:"logic_group_type,omitempty"`
    

}
