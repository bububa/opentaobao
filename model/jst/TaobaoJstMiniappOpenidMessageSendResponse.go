package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单个openId用户短信发送 APIResponse
taobao.jst.miniapp.openid.message.send

单个openId用户短信发送
*/
type TaobaoJstMiniappOpenidMessageSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstMiniappOpenidMessageSendResponse `json:"taobao_jst_miniapp_openid_message_send_response,omitempty"`
}

type TaobaoJstMiniappOpenidMessageSendResponse struct {

    // 短信回执码
    Result   string `json:"result,omitempty"`

    // 请求code
    RCode   int64 `json:"r_code,omitempty"`

    // 请求失败错误信息
    RErrMsg   string `json:"r_err_msg,omitempty"`

}
