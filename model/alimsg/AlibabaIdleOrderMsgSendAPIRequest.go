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

// NewAlibabaIdleOrderMsgSendRequest 初始化AlibabaIdleOrderMsgSendAPIRequest对象
func NewAlibabaIdleOrderMsgSendRequest() *AlibabaIdleOrderMsgSendAPIRequest {
	return &AlibabaIdleOrderMsgSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleOrderMsgSendAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.order.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleOrderMsgSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaIdleOrderMsgSendAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaIdleOrderMsgSendAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Text Setter
// 消息发送内容
func (r *AlibabaIdleOrderMsgSendAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// Get Text Getter
func (r AlibabaIdleOrderMsgSendAPIRequest) GetText() string {
	return r._text
}
