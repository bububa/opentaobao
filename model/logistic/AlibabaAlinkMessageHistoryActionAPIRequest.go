package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkmessagehistoryactionAPIRequest 操作历史消息 API请求
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
type AlibabaalinkmessagehistoryactionAPIRequest struct {
	model.Params
	// 消息id
	_index string
	// 删除：delete,已读：read
	_action string
}

// NewAlibabaalinkmessagehistoryactionRequest 初始化AlibabaalinkmessagehistoryactionAPIRequest对象
func NewAlibabaalinkmessagehistoryactionRequest() *AlibabaalinkmessagehistoryactionAPIRequest {
	return &AlibabaalinkmessagehistoryactionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkmessagehistoryactionAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.history.action"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkmessagehistoryactionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkmessagehistoryactionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIndex is Index Setter
// 消息id
func (r *AlibabaalinkmessagehistoryactionAPIRequest) SetIndex(_index string) error {
	r._index = _index
	r.Set("index", _index)
	return nil
}

// GetIndex Index Getter
func (r AlibabaalinkmessagehistoryactionAPIRequest) GetIndex() string {
	return r._index
}

// SetAction is Action Setter
// 删除：delete,已读：read
func (r *AlibabaalinkmessagehistoryactionAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaalinkmessagehistoryactionAPIRequest) GetAction() string {
	return r._action
}
