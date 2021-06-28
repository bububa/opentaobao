package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个openId用户短信发送 APIResponse
taobao.jst.miniapp.openid.message.send

单个openId用户短信发送
*/
type TaobaoJstMiniappOpenidMessageSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_miniapp_openid_message_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 短信回执码
    
    Result   string `json:"result,omitempty" xml:"