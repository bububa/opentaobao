package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
跟进客资状态接口 APIResponse
alibaba.lsy.crm.update

同步客资状态接口
*/
type AlibabaLsyCrmUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLsyCrmUpdateResponse `json:"alibaba_lsy_crm_update_response,omitempty"` 
    AlibabaLsyCrmUpdateResponse
}

/* model for simplify = false
type AlibabaLsyCrmUpdateResponse struct {

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty"`
    

    // 错误码
    
    FailCode   string `json:"fail_code,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

    // 是否成功跟进客资状态
    
    Data   bool `json:"data,omitempty"`
    

}
*/

type AlibabaLsyCrmUpdateResponse struct {

    // 错误提示
    FailMsg   string `json:"fail_msg,omitempty"`

    // 错误码
    FailCode   string `json:"fail_code,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

    // 是否成功跟进客资状态
    Data   bool `json:"data,omitempty"`

}
