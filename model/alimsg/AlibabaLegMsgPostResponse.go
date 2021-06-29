package alimsg

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
集团法务消息发送 APIResponse
alibaba.leg.msg.post

消息发送能力
*/
type AlibabaLegMsgPostAPIResponse struct {
    model.CommonResponse
    AlibabaLegMsgPostResponse
}

type AlibabaLegMsgPostResponse struct {
    XMLName xml.Name `xml:"alibaba_leg_msg_post_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
