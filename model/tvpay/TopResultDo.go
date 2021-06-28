package tvpay

// TopResultDo 
/* model for simplify = false
type TopResultDo struct {

    // 状态码
    
    Code   string `json:"code,omitempty"`
    

    // 结构体
    
    Data  *struct {
        GetLoginInfoByOrderResultDo  *GetLoginInfoByOrderResultDo `json:"get_login_info_by_order_result_do,omitempty"`
    } `json:"data,omitempty"`
    

    // 消息
    
    Message   string `json:"message,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TopResultDo 
type TopResultDo struct {

    // 状态码
    Code   string `json:"code,omitempty"`

    // 结构体
    Data   *GetLoginInfoByOrderResultDo `json:"data,omitempty"`

    // 消息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
