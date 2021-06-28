package tmallhk

// DataResult 
/* model for simplify = false
type DataResult struct {

    // 参数code
    
    Code   string `json:"code,omitempty"`
    

    // 参数msg
    
    Msg   string `json:"msg,omitempty"`
    

    // 是否正常
    
    Success   bool `json:"success,omitempty"`
    

    // obj
    
    Obj  *struct {
        CCICCheckCodeDO  *CCICCheckCodeDO `json:"ccic_check_code_do,omitempty"`
    } `json:"obj,omitempty"`
    

}
*/

// DataResult 
type DataResult struct {

    // 参数code
    Code   string `json:"code,omitempty"`

    // 参数msg
    Msg   string `json:"msg,omitempty"`

    // 是否正常
    Success   bool `json:"success,omitempty"`

    // obj
    Obj   *CCICCheckCodeDO `json:"obj,omitempty"`

}
