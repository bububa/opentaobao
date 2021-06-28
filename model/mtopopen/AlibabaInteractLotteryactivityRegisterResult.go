package mtopopen

// AlibabaInteractLotteryactivityRegisterResult 
/* model for simplify = false
type AlibabaInteractLotteryactivityRegisterResult struct {

    // 返回的数据
    
    Data  *struct {
        ActivityLotteryWriteResult  *ActivityLotteryWriteResult `json:"activity_lottery_write_result,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // xxx
    
    ModuleMap  *struct {
        Modulemap  *Modulemap `json:"modulemap,omitempty"`
    } `json:"module_map,omitempty"`
    

    // 注册抽奖活动失败
    
    Success   bool `json:"success,omitempty"`
    

    // 方便追踪请求的唯一标识
    
    TraceId   string `json:"trace_id,omitempty"`
    

}
*/

// AlibabaInteractLotteryactivityRegisterResult 
type AlibabaInteractLotteryactivityRegisterResult struct {

    // 返回的数据
    Data   *ActivityLotteryWriteResult `json:"data,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // xxx
    ModuleMap   *Modulemap `json:"module_map,omitempty"`

    // 注册抽奖活动失败
    Success   bool `json:"success,omitempty"`

    // 方便追踪请求的唯一标识
    TraceId   string `json:"trace_id,omitempty"`

}
