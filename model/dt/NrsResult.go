package dt

// NrsResult 
/* model for simplify = false
type NrsResult struct {

    // 返回数据
    
    Data  *struct {
        RecongnizeItemInfo  *RecongnizeItemInfo `json:"recongnize_item_info,omitempty"`
    } `json:"data,omitempty"`
    

    // 接口调用标志
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

}
*/

// NrsResult 
type NrsResult struct {

    // 返回数据
    Data   *RecongnizeItemInfo `json:"data,omitempty"`

    // 接口调用标志
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

}
