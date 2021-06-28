package logistic

// ReverseEventInfoDto 
/* model for simplify = false
type ReverseEventInfoDto struct {

    // 销退单ID
    
    Id   string `json:"id,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 扩展字段
    
    Extra   string `json:"extra,omitempty"`
    

    // 值
    
    Value   string `json:"value,omitempty"`
    

    // 类型(1=销退单状态变更)
    
    Type   int64 `json:"type,omitempty"`
    

}
*/

// ReverseEventInfoDto 
type ReverseEventInfoDto struct {

    // 销退单ID
    Id   string `json:"id,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 扩展字段
    Extra   string `json:"extra,omitempty"`

    // 值
    Value   string `json:"value,omitempty"`

    // 类型(1=销退单状态变更)
    Type   int64 `json:"type,omitempty"`

}
