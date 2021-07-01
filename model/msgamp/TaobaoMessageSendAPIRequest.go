package msgamp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMessageSendAPIRequest
消息发送 API请求
taobao.message.send

消息发送接口 */
type TaobaoMessageSendAPIRequest struct {
	model.Params
	// 消息发送相关参数
	_sendMessageReq *SendMessageReq
}

// New
