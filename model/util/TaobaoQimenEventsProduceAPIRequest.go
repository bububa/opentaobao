package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenEventsProduceAPIRequest) Reset() {
	r._messages = r._messages[:0]
	r.Params.ToZero()
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

var poolTaobaoQimenEventsProduceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenEventsProduceRequest()
	},
}

// GetTaobaoQimenEventsProduceRequest 从 sync.Pool 获取 TaobaoQimenEventsProduceAPIRequest
func GetTaobaoQimenEventsProduceAPIRequest() *TaobaoQimenEventsProduceAPIRequest {
	return poolTaobaoQimenEventsProduceAPIRequest.Get().(*TaobaoQimenEventsProduceAPIRequest)
}

// ReleaseTaobaoQimenEventsProduceAPIRequest 将 TaobaoQimenEventsProduceAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenEventsProduceAPIRequest(v *TaobaoQimenEventsProduceAPIRequest) {
	v.Reset()
	poolTaobaoQimenEventsProduceAPIRequest.Put(v)
}
