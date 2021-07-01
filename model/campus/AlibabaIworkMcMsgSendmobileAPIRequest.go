package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIworkMcMsgSendmobileAPIRequest
发送消息给手机用户 API请求
alibaba.iwork.mc.msg.sendmobile

给手机用户发送对应操作结果的消息 */
type AlibabaIworkMcMsgSendmobileAPIRequest struct {
	model.Params
	// 消息对象
	_mobileReceiverMessageEvent *MobileReceiverMessageEvent
}

// New
