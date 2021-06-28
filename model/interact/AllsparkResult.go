package interact

// AllsparkResult 
/* model for simplify = false
type AllsparkResult struct {

    // 出错提示
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 出错代码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 服务结果对象
    
    Data  *struct {
        ActivityWriteResult  *ActivityWriteResult `json:"activity_write_result,omitempty"`
    } `json:"data,omitempty"`
    

    // 是否注册成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AllsparkResult 
type AllsparkResult struct {

    // 出错提示
    ErrMsg   string `json:"err_msg,omitempty"`

    // 出错代码
    ErrCode   string `json:"err_code,omitempty"`

    // 服务结果对象
    Data   *ActivityWriteResult `json:"data,omitempty"`

    // 是否注册成功
    Success   bool `json:"success,omitempty"`

}
