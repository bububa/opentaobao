package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序消息推送 APIResponse
alibaba.lsy.miniapp.msg.push

零售云小程序消息推送，推送消息至零售云（喵零等）
*/
type AlibabaLsyMiniappMsgPushAPIResponse struct {
    model.CommonResponse
    AlibabaLsyMiniappMsgPushResponse
}

type AlibabaLsyMiniappMsgPushResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_miniapp_msg_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误提示
    
    FailMsg   string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`

    
    // 错误码
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
}
