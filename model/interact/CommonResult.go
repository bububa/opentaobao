package interact

// CommonResult 
/* model for simplify = false
type CommonResult struct {

    // 是否调用成功
    
    Succ   bool `json:"succ,omitempty"`
    

    // 错误码
    
    Code   string `json:"code,omitempty"`
    

    // 错误描述
    
    Msg   string `json:"msg,omitempty"`
    

    // 创建的互动实例
    
    Data  *struct {
        InteractDTO  *InteractDTO `json:"interact_dto,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// CommonResult 
type CommonResult struct {

    // 是否调用成功
    Succ   bool `json:"succ,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

    // 错误描述
    Msg   string `json:"msg,omitempty"`

    // 创建的互动实例
    Data   *InteractDTO `json:"data,omitempty"`

}
