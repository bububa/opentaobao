package wdk

// AlibabaWdkElemeBillDetailGetApiResult 
/* model for simplify = false
type AlibabaWdkElemeBillDetailGetApiResult struct {

    // 账单信息
    
    Model  *struct {
        EleBillBo  *EleBillBo `json:"ele_bill_bo,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误描述
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkElemeBillDetailGetApiResult 
type AlibabaWdkElemeBillDetailGetApiResult struct {

    // 账单信息
    Model   *EleBillBo `json:"model,omitempty"`

    // 错误描述
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

}
