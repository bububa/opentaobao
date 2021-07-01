package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送自定义openim消息 API返回值 
taobao.openim.custmsg.push

isv通过该接口给openim用户推送自定义消息
*/
type TaobaoOpenimCustmsgPushAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimCustmsgPushAPIResponseModel
}

// 推送自定义openim消息 成功返回结果
type TaobaoOpenimCustmsgPushAPIResponseModel struct {
    XMLName xml.Name `xml:"openim_custmsg_push_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 消息id，用于定位问题
    Msgid   int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
}
