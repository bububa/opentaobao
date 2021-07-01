package alimsg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleOrderMsgSendAPIRequest
虚拟发货消息发送接口 API请求
alibaba.idle.order.msg.send

用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景 */
type AlibabaIdleOrderMsgSendAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 消息发送内容
	_text string
}

// New
