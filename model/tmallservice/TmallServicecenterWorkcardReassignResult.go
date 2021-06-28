package tmallservice

// TmallServicecenterWorkcardReassignResult 
/* model for simplify = false
type TmallServicecenterWorkcardReassignResult struct {

    // 调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 改派后的新的工单id
    
    ResultData   int64 `json:"result_data,omitempty"`
    

}
*/

// TmallServicecenterWorkcardReassignResult 
type TmallServicecenterWorkcardReassignResult struct {

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 改派后的新的工单id
    ResultData   int64 `json:"result_data,omitempty"`

}
