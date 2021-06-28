package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openim标准消息发送 APIResponse
taobao.openim.immsg.push

服务端对openim用户发送标准消息，包括文字、语音、图片等。
*/
type TaobaoOpenimImmsgPushAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimImmsgPushResponse
}

type TaobaoOpenimImmsgPushResponse struct {
    XMLName xml.Name `xml:"openim_immsg_push_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息id，方便定位问题
    
    Msgid   int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`

    
}
