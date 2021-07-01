package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIworkMcMsgSenddefaultAPIRequest
给注册用户发送消息 API请求
alibaba.iwork.mc.msg.senddefault

给神鲸注册用户发送对应操作结果的消息 */
type AlibabaIworkMcMsgSenddefaultAPIRequest struct {
	model.Params
	// 消息对象
	_messageEvent *DefaultMessageEvent
}

// New
