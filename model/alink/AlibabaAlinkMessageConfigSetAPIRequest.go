package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageConfigSetAPIRequest 消息提醒开关 API请求
// alibaba.alink.message.config.set
//
// 阿里智能消息开关
type AlibabaAlinkMessageConfigSetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 0：开启，1：关闭
	_pushDisabled string
}

// NewAlibabaAlinkMessageConfigSetRequest 初始化AlibabaAlinkMessageConfigSetAPIRequest对象
func NewAlibabaAlinkMessageConfigSetRequest() *AlibabaAlinkMessageConfigSetAPIRequest {
	return &AlibabaAlinkMessageConfigSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.config.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Uuid Setter
// 设备id
func (r *AlibabaAlinkMessageConfigSetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is PushDisabled Setter
// 0：开启，1：关闭
func (r *AlibabaAlinkMessageConfigSetAPIRequest) SetPushDisabled(_pushDisabled string) error {
	r._pushDisabled = _pushDisabled
	r.Set("push_disabled", _pushDisabled)
	return nil
}

// Get PushDisabled Getter
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetPushDisabled() string {
	return r._pushDisabled
}
