package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopconnectoreventpublishAPIRequest 连接器事件发布 API请求
// taobao.top.connector.event.publish
//
// 连接器事件发布
type TaobaotopconnectoreventpublishAPIRequest struct {
	model.Params
	// 发布事件列表
	_entryList *EventPublishThirdPartyEntry
}

// NewTaobaotopconnectoreventpublishRequest 初始化TaobaotopconnectoreventpublishAPIRequest对象
func NewTaobaotopconnectoreventpublishRequest() *TaobaotopconnectoreventpublishAPIRequest {
	return &TaobaotopconnectoreventpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopconnectoreventpublishAPIRequest) GetApiMethodName() string {
	return "taobao.top.connector.event.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopconnectoreventpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopconnectoreventpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntryList is EntryList Setter
// 发布事件列表
func (r *TaobaotopconnectoreventpublishAPIRequest) SetEntryList(_entryList *EventPublishThirdPartyEntry) error {
	r._entryList = _entryList
	r.Set("entry_list", _entryList)
	return nil
}

// GetEntryList EntryList Getter
func (r TaobaotopconnectoreventpublishAPIRequest) GetEntryList() *EventPublishThirdPartyEntry {
	return r._entryList
}
