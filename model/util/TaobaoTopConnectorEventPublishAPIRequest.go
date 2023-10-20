package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopConnectorEventPublishAPIRequest 连接器事件发布 API请求
// taobao.top.connector.event.publish
//
// 连接器事件发布
type TaobaoTopConnectorEventPublishAPIRequest struct {
	model.Params
	// 发布事件列表
	_entryList *EventPublishThirdPartyEntry
}

// NewTaobaoTopConnectorEventPublishRequest 初始化TaobaoTopConnectorEventPublishAPIRequest对象
func NewTaobaoTopConnectorEventPublishRequest() *TaobaoTopConnectorEventPublishAPIRequest {
	return &TaobaoTopConnectorEventPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopConnectorEventPublishAPIRequest) Reset() {
	r._entryList = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopConnectorEventPublishAPIRequest) GetApiMethodName() string {
	return "taobao.top.connector.event.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopConnectorEventPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopConnectorEventPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntryList is EntryList Setter
// 发布事件列表
func (r *TaobaoTopConnectorEventPublishAPIRequest) SetEntryList(_entryList *EventPublishThirdPartyEntry) error {
	r._entryList = _entryList
	r.Set("entry_list", _entryList)
	return nil
}

// GetEntryList EntryList Getter
func (r TaobaoTopConnectorEventPublishAPIRequest) GetEntryList() *EventPublishThirdPartyEntry {
	return r._entryList
}

var poolTaobaoTopConnectorEventPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopConnectorEventPublishRequest()
	},
}

// GetTaobaoTopConnectorEventPublishRequest 从 sync.Pool 获取 TaobaoTopConnectorEventPublishAPIRequest
func GetTaobaoTopConnectorEventPublishAPIRequest() *TaobaoTopConnectorEventPublishAPIRequest {
	return poolTaobaoTopConnectorEventPublishAPIRequest.Get().(*TaobaoTopConnectorEventPublishAPIRequest)
}

// ReleaseTaobaoTopConnectorEventPublishAPIRequest 将 TaobaoTopConnectorEventPublishAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopConnectorEventPublishAPIRequest(v *TaobaoTopConnectorEventPublishAPIRequest) {
	v.Reset()
	poolTaobaoTopConnectorEventPublishAPIRequest.Put(v)
}
