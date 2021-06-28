package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序消息推送 APIResponse
alibaba.lsy.miniapp.msg.push

零售云小程序消息推送，推送消息至零售云（喵零等）
*/
type AlibabaLsyMiniappMsgPushAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLsyMiniappMsgPushResponse `json:"alibaba_lsy_miniapp_msg_push_response,omitempty"` 
    AlibabaLsyMiniappMsgPushResponse
}

/* model for simplify = false
type AlibabaLsyMiniappMsgPushResponse struct {

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty"`
    

    // 错误码
    
    FailCode   string `json:"fail_code,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

}
*/

type AlibabaLsyMiniappMsgPushResponse struct {

    // 错误提示
    FailMsg   string `json:"fail_msg,omitempty"`

    // 错误码
    FailCode   string `json:"fail_code,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

}
