package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageHistoryActionAPIRequest 操作历史消息 API请求
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
type AlibabaAlinkMessageHistoryActionAPIRequest struct {
	model.Params
	// 消息id
	_index string
	// 删除：delete,已读：read
	_action string
}

// NewAlibabaAlinkMessageHistoryActionRequest 初始化AlibabaAlinkMessageHistoryActionAPIRequest对象
func NewAlibabaAlinkMessageHistoryActionRequest() *AlibabaAlinkMessageHistoryActionAPIRequest {
	return &AlibabaAlinkMessageHistoryActionAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlinkMessageHistoryActionAPIRequest) Reset() {
	r._index = ""
	r._action = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.history.action"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIndex is Index Setter
// 消息id
func (r *AlibabaAlinkMessageHistoryActionAPIRequest) SetIndex(_index string) error {
	r._index = _index
	r.Set("index", _index)
	return nil
}

// GetIndex Index Getter
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetIndex() string {
	return r._index
}

// SetAction is Action Setter
// 删除：delete,已读：read
func (r *AlibabaAlinkMessageHistoryActionAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetAction() string {
	return r._action
}

var poolAlibabaAlinkMessageHistoryActionAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlinkMessageHistoryActionRequest()
	},
}

// GetAlibabaAlinkMessageHistoryActionRequest 从 sync.Pool 获取 AlibabaAlinkMessageHistoryActionAPIRequest
func GetAlibabaAlinkMessageHistoryActionAPIRequest() *AlibabaAlinkMessageHistoryActionAPIRequest {
	return poolAlibabaAlinkMessageHistoryActionAPIRequest.Get().(*AlibabaAlinkMessageHistoryActionAPIRequest)
}

// ReleaseAlibabaAlinkMessageHistoryActionAPIRequest 将 AlibabaAlinkMessageHistoryActionAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlinkMessageHistoryActionAPIRequest(v *AlibabaAlinkMessageHistoryActionAPIRequest) {
	v.Reset()
	poolAlibabaAlinkMessageHistoryActionAPIRequest.Put(v)
}
