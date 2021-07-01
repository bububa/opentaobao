package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个openId用户短信发送 API返回值 
taobao.jst.miniapp.openid.message.send

单个openId用户短信发送
*/
type TaobaoJstMiniappOpenidMessageSendAPIResponse struct {
    model.CommonResponse
    TaobaoJstMiniappOpenidMessageSendAPIResponseModel
}

// 单个openId用户短信发送 成功返回结果
type TaobaoJstMiniappOpenidMessageSendAPIResponseModel struct {
    XMLName xml.Name `xml:"jst_miniapp_openid_message_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 短信回执码
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 请求code
    RCode   int64 `json:"r_code,omitempty" xml:"r_code,omitempty"`
    // 请求失败错误信息
    RErrMsg   string `json:"r_err_msg,omitempty" xml:"r_err_msg,omitempty"`
}
