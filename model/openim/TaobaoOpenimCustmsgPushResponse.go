package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送自定义openim消息 APIResponse
taobao.openim.custmsg.push

isv通过该接口给openim用户推送自定义消息
*/
type TaobaoOpenimCustmsgPushAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimCustmsgPushResponse
}

type TaobaoOpenimCustmsgPushResponse struct {
    XMLName xml.Name `xml:"openim_custmsg_push_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息id，用于定位问题
    
    Msgid   int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`

    
}
