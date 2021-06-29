package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔数据paas短信发送接口 API返回值 
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
*/
type TaobaoJstSmsMessageSendAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsMessageSendResponse
}

// 聚石塔数据paas短信发送接口 成功返回结果
type TaobaoJstSmsMessageSendResponse struct {
    XMLName xml.Name `xml:"jst_sms_message_send_response"`
    // 参数错误
    RequestCode   string `json:"request_code,omitempty" xml:"request_code,omitempty"`
    // 请求成功
    RequestSuccess   bool `json:"request_success,omitempty" xml:"request_success,omitempty"`
    // 1234
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 空
    Module   string `json:"module,omitempty" xml:"module,omitempty"`
    // 参数错误
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
