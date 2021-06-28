package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔数据paas短信发送接口 APIResponse
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
*/
type TaobaoJstSmsMessageSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_message_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 参数错误
    
    RequestCode   string `json:"request_code,omitempty" xml:"