package tmallservice

// AlibabaServicecenterSpserviceorderQueryResult 
/* model for simplify = false
type AlibabaServicecenterSpserviceorderQueryResult struct {

    // 分页数据
    
    ResultData  *struct {
        Paged  *Paged `json:"paged,omitempty"`
    } `json:"result_data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误描述
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

}
*/

// AlibabaServicecenterSpserviceorderQueryResult 
type AlibabaServicecenterSpserviceorderQueryResult struct {

    // 分页数据
    ResultData   *Paged `json:"result_data,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误描述
    MsgInfo   string `json:"msg_info,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

}
