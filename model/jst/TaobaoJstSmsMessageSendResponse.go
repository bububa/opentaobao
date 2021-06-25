package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔数据paas短信发送接口 APIResponse
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
*/
type TaobaoJstSmsMessageSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsMessageSendResponse `json:"taobao_jst_sms_message_send_response,omitempty"`
}

type TaobaoJstSmsMessageSendResponse struct {

    // 参数错误
    RequestCode   string `json:"request_code,omitempty"`

    // 请求成功
    RequestSuccess   bool `json:"request_success,omitempty"`

    // 1234
    RequestId   string `json:"request_id,omitempty"`

    // 空
    Module   string `json:"module,omitempty"`

    // 参数错误
    Message   string `json:"message,omitempty"`

}
