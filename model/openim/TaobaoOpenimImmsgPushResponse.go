package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openim标准消息发送 APIResponse
taobao.openim.immsg.push

服务端对openim用户发送标准消息，包括文字、语音、图片等。
*/
type TaobaoOpenimImmsgPushAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimImmsgPushResponse `json:"taobao_openim_immsg_push_response,omitempty"`
}

type TaobaoOpenimImmsgPushResponse struct {

    // 消息id，方便定位问题
    Msgid   int64 `json:"msgid,omitempty"`

}
