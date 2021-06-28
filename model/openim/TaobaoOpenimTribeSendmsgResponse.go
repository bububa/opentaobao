package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发送群消息 APIResponse
taobao.openim.tribe.sendmsg

发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
*/
type TaobaoOpenimTribeSendmsgAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeSendmsgResponse `json:"openim_tribe_sendmsg_response,omitempty"` 
    TaobaoOpenimTribeSendmsgResponse
}

/* model for simplify = false
type TaobaoOpenimTribeSendmsgResponse struct {

    // 错误码
    
    TribeCode   int64 `json:"tribe_code,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

type TaobaoOpenimTribeSendmsgResponse struct {

    // 错误码
    TribeCode   int64 `json:"tribe_code,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

}
