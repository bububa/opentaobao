package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopEventPublishAPIRequest 同步事件发布 API请求
// taobao.top.event.publish
//
// 同步事件发布
type TaobaoTopEventPublishAPIRequest struct {
	model.Params
	// 事件编码
	_triggerCode string
	// 事件内容
	_content string
}

// NewTaobaoTopEventPublishRequest 初始化TaobaoTopEventPublishAPIRequest对象
func NewTaobaoTopEventPublishRequest() *TaobaoTopEventPublishAPIRequest {
	return &TaobaoTopEventPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopEventPublishAPIRequest) GetApiMethodName() string {
	return "taobao.top.event.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopEventPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopEventPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTriggerCode is TriggerCode Setter
// 事件编码
func (r *TaobaoTopEventPublishAPIRequest) SetTriggerCode(_triggerCode string) error {
	r._triggerCode = _triggerCode
	r.Set("trigger_code", _triggerCode)
	return nil
}

// GetTriggerCode TriggerCode Getter
func (r TaobaoTopEventPublishAPIRequest) GetTriggerCode() string {
	return r._triggerCode
}

// SetContent is Content Setter
// 事件内容
func (r *TaobaoTopEventPublishAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoTopEventPublishAPIRequest) GetContent() string {
	return r._content
}
