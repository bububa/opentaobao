package alimsg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleordermsgsendAPIRequest 虚拟发货消息发送接口 API请求
// alibaba.idle.order.msg.send
//
// 用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
type AlibabaidleordermsgsendAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 消息发送内容
	_text string
}

// NewAlibabaidleordermsgsendRequest 初始化AlibabaidleordermsgsendAPIRequest对象
func NewAlibabaidleordermsgsendRequest() *AlibabaidleordermsgsendAPIRequest {
	return &AlibabaidleordermsgsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleordermsgsendAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.order.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleordermsgsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleordermsgsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaidleordermsgsendAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaidleordermsgsendAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetText is Text Setter
// 消息发送内容
func (r *AlibabaidleordermsgsendAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r AlibabaidleordermsgsendAPIRequest) GetText() string {
	return r._text
}
