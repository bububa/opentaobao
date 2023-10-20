package alink

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlinkMessageConfigSetAPIRequest) Reset() {
	r._uuid = ""
	r._pushDisabled = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.config.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkMessageConfigSetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPushDisabled is PushDisabled Setter
// 0：开启，1：关闭
func (r *AlibabaAlinkMessageConfigSetAPIRequest) SetPushDisabled(_pushDisabled string) error {
	r._pushDisabled = _pushDisabled
	r.Set("push_disabled", _pushDisabled)
	return nil
}

// GetPushDisabled PushDisabled Getter
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetPushDisabled() string {
	return r._pushDisabled
}

var poolAlibabaAlinkMessageConfigSetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlinkMessageConfigSetRequest()
	},
}

// GetAlibabaAlinkMessageConfigSetRequest 从 sync.Pool 获取 AlibabaAlinkMessageConfigSetAPIRequest
func GetAlibabaAlinkMessageConfigSetAPIRequest() *AlibabaAlinkMessageConfigSetAPIRequest {
	return poolAlibabaAlinkMessageConfigSetAPIRequest.Get().(*AlibabaAlinkMessageConfigSetAPIRequest)
}

// ReleaseAlibabaAlinkMessageConfigSetAPIRequest 将 AlibabaAlinkMessageConfigSetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlinkMessageConfigSetAPIRequest(v *AlibabaAlinkMessageConfigSetAPIRequest) {
	v.Reset()
	poolAlibabaAlinkMessageConfigSetAPIRequest.Put(v)
}
