package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimCustmsgPushAPIRequest
推送自定义openim消息 API请求
taobao.openim.custmsg.push

isv通过该接口给openim用户推送自定义消息 */
type TaobaoOpenimCustmsgPushAPIRequest struct {
	model.Params
	// 自定义消息内容
	_custmsg *CustMsg
}

// New
