package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopeventpublishAPIRequest 同步事件发布 API请求
// taobao.top.event.publish
//
// 同步事件发布
type TaobaotopeventpublishAPIRequest struct {
	model.Params
	// 事件编码
	_triggerCode string
	// 事件内容
	_content string
}

// NewTaobaotopeventpublishRequest 初始化TaobaotopeventpublishAPIRequest对象
func NewTaobaotopeventpublishRequest() *TaobaotopeventpublishAPIRequest {
	return &TaobaotopeventpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopeventpublishAPIRequest) GetApiMethodName() string {
	return "taobao.top.event.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopeventpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopeventpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTriggerCode is TriggerCode Setter
// 事件编码
func (r *TaobaotopeventpublishAPIRequest) SetTriggerCode(_triggerCode string) error {
	r._triggerCode = _triggerCode
	r.Set("trigger_code", _triggerCode)
	return nil
}

// GetTriggerCode TriggerCode Getter
func (r TaobaotopeventpublishAPIRequest) GetTriggerCode() string {
	return r._triggerCode
}

// SetContent is Content Setter
// 事件内容
func (r *TaobaotopeventpublishAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaotopeventpublishAPIRequest) GetContent() string {
	return r._content
}
