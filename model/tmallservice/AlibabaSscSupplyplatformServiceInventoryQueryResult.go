package tmallservice

// AlibabaSscSupplyplatformServiceInventoryQueryResult 
/* model for simplify = false
type AlibabaSscSupplyplatformServiceInventoryQueryResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // json字符串。key是时间片，格式为yyyy-MM-dd或yyyy-MM-dd HH:mm-HH:mm，value是当前库存值
    
    Value   string `json:"value,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 对外展示的错误信息
    
    DisplayMsg   string `json:"display_msg,omitempty"`
    

}
*/

// AlibabaSscSupplyplatformServiceInventoryQueryResult 
type AlibabaSscSupplyplatformServiceInventoryQueryResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // json字符串。key是时间片，格式为yyyy-MM-dd或yyyy-MM-dd HH:mm-HH:mm，value是当前库存值
    Value   string `json:"value,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 对外展示的错误信息
    DisplayMsg   string `json:"display_msg,omitempty"`

}
