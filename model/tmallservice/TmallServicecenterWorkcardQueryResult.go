package tmallservice

// TmallServicecenterWorkcardQueryResult 
/* model for simplify = false
type TmallServicecenterWorkcardQueryResult struct {

    // 分页数据
    
    ResultData  *struct {
        Paged  *Paged `json:"paged,omitempty"`
    } `json:"result_data,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServicecenterWorkcardQueryResult 
type TmallServicecenterWorkcardQueryResult struct {

    // 分页数据
    ResultData   *Paged `json:"result_data,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
