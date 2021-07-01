package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkMessageHistoryActionAPIRequest
操作历史消息 API请求
alibaba.alink.message.history.action

阿里智能操作历史消息 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.history.action"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Index Setter
// 消息id
func (r *AlibabaAlinkMessageHistoryActionAPIRequest) SetIndex(_index string) error {
	r._index = _index
	r.Set("index", _index)
	return nil
}

// Get Index Getter
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetIndex() string {
	return r._index
}

// Set is Action Setter
// 删除：delete,已读：read
func (r *AlibabaAlinkMessageHistoryActionAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// Get Action Getter
func (r AlibabaAlinkMessageHistoryActionAPIRequest) GetAction() string {
	return r._action
}
