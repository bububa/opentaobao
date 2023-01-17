package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEventsProduceAPIRequest 批量发送奇门事件 API请求
// taobao.qimen.events.produce
//
// 批量发送消息
type TaobaoQimenEventsProduceAPIRequest struct {
	model.Params
	// 奇门事件列表, 最多50条
	_messages []QimenEvent
}

// NewTaobaoQimenEventsProduceRequest 初始化TaobaoQimenEventsProduceAPIRequest对象
func NewTaobaoQimenEventsProduceRequest() *TaobaoQimenEventsProduceAPIRequest {
	return &TaobaoQimenEventsProduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenEventsProduceAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.events.produce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenEventsProduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenEventsProduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessages is Messages Setter
// 奇门事件列表, 最多50条
func (r *TaobaoQimenEventsProduceAPIRequest) SetMessages(_messages []QimenEvent) error {
	r._messages = _messages
	r.Set("messages", _messages)
	return nil
}

// GetMessages Messages Getter
func (r TaobaoQimenEventsProduceAPIRequest) GetMessages() []QimenEvent {
	return r._messages
}
