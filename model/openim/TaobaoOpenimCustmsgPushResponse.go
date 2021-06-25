package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推送自定义openim消息 APIResponse
taobao.openim.custmsg.push

isv通过该接口给openim用户推送自定义消息
*/
type TaobaoOpenimCustmsgPushAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimCustmsgPushResponse `json:"taobao_openim_custmsg_push_response,omitempty"`
}

type TaobaoOpenimCustmsgPushResponse struct {

    // 消息id，用于定位问题
    Msgid   int64 `json:"msgid,omitempty"`

}
