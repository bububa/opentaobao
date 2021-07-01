package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmMemberMessageHandleAPIRequest
处理Usc会员消息接口 API请求
alibaba.westcrm.member.message.handle

处理Usc会员消息接口 */
type AlibabaWestcrmMemberMessageHandleAPIRequest struct {
	model.Params
	// 消息类型
	_messageType string
	// 消息内容
	_messageContent string
}

// New
