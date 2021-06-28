package aliqin

// AlibabaAliqinFcIotCardStatusResult 
/* model for simplify = false
type AlibabaAliqinFcIotCardStatusResult struct {

    // "MsisdnStatus":"卡状态","MSISDN":"MSISDN号","ICCID":"ICCID号"
    
    Model   string `json:"model,omitempty"`
    

    // 1.成功；2.失败
    
    Code   string `json:"code,omitempty"`
    

    // 状态
    
    Success   bool `json:"success,omitempty"`
    

    // 错误信息
    
    Msg   string `json:"msg,omitempty"`
    

}
*/

// AlibabaAliqinFcIotCardStatusResult 
type AlibabaAliqinFcIotCardStatusResult struct {

    // "MsisdnStatus":"卡状态","MSISDN":"MSISDN号","ICCID":"ICCID号"
    Model   string `json:"model,omitempty"`

    // 1.成功；2.失败
    Code   string `json:"code,omitempty"`

    // 状态
    Success   bool `json:"success,omitempty"`

    // 错误信息
    Msg   string `json:"msg,omitempty"`

}
