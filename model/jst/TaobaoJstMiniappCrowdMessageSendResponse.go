package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动短信发送 APIResponse
taobao.jst.miniapp.crowd.message.send

小程序活动短信发送
*/
type TaobaoJstMiniappCrowdMessageSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstMiniappCrowdMessageSendResponse `json:"jst_miniapp_crowd_message_send_response,omitempty"` 
    TaobaoJstMiniappCrowdMessageSendResponse
}

/* model for simplify = false
type TaobaoJstMiniappCrowdMessageSendResponse struct {

    // 发送的短信条数
    
    Result   int64 `json:"result,omitempty"`
    

    // 发送成功
    
    RCode   int64 `json:"r_code,omitempty"`
    

}
*/

type TaobaoJstMiniappCrowdMessageSendResponse struct {

    // 发送的短信条数
    Result   int64 `json:"result,omitempty"`

    // 发送成功
    RCode   int64 `json:"r_code,omitempty"`

}
