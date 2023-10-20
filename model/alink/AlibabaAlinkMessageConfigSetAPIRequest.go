package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkmessageconfigsetAPIRequest 消息提醒开关 API请求
// alibaba.alink.message.config.set
//
// 阿里智能消息开关
type AlibabaalinkmessageconfigsetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 0：开启，1：关闭
	_pushDisabled string
}

// NewAlibabaalinkmessageconfigsetRequest 初始化AlibabaalinkmessageconfigsetAPIRequest对象
func NewAlibabaalinkmessageconfigsetRequest() *AlibabaalinkmessageconfigsetAPIRequest {
	return &AlibabaalinkmessageconfigsetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkmessageconfigsetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.config.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkmessageconfigsetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkmessageconfigsetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkmessageconfigsetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkmessageconfigsetAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPushDisabled is PushDisabled Setter
// 0：开启，1：关闭
func (r *AlibabaalinkmessageconfigsetAPIRequest) SetPushDisabled(_pushDisabled string) error {
	r._pushDisabled = _pushDisabled
	r.Set("push_disabled", _pushDisabled)
	return nil
}

// GetPushDisabled PushDisabled Getter
func (r AlibabaalinkmessageconfigsetAPIRequest) GetPushDisabled() string {
	return r._pushDisabled
}
