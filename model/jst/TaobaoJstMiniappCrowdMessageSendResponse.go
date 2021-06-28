package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动短信发送 APIResponse
taobao.jst.miniapp.crowd.message.send

小程序活动短信发送
*/
type TaobaoJstMiniappCrowdMessageSendAPIResponse struct {
    model.CommonResponse
    TaobaoJstMiniappCrowdMessageSendResponse
}

type TaobaoJstMiniappCrowdMessageSendResponse struct {
    XMLName xml.Name `xml:"jst_miniapp_crowd_message_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 发送的短信条数
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
    // 发送成功
    
    RCode   int64 `json:"r_code,omitempty" xml:"r_code,omitempty"`

    
}
