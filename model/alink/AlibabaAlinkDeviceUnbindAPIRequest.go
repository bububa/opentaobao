package alink

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceUnbindAPIRequest 解绑设备 API请求
// alibaba.alink.device.unbind
//
// 阿里智能解绑设备
type AlibabaAlinkDeviceUnbindAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaAlinkDeviceUnbindRequest 初始化AlibabaAlinkDeviceUnbindAPIRequest对象
func NewAlibabaAlinkDeviceUnbindRequest() *AlibabaAlinkDeviceUnbindAPIRequest {
	return &AlibabaAlinkDeviceUnbindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlinkDeviceUnbindAPIRequest) Reset() {
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceUnbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetUuid() string {
	return r._uuid
}

var poolAlibabaAlinkDeviceUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlinkDeviceUnbindRequest()
	},
}

// GetAlibabaAlinkDeviceUnbindRequest 从 sync.Pool 获取 AlibabaAlinkDeviceUnbindAPIRequest
func GetAlibabaAlinkDeviceUnbindAPIRequest() *AlibabaAlinkDeviceUnbindAPIRequest {
	return poolAlibabaAlinkDeviceUnbindAPIRequest.Get().(*AlibabaAlinkDeviceUnbindAPIRequest)
}

// ReleaseAlibabaAlinkDeviceUnbindAPIRequest 将 AlibabaAlinkDeviceUnbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlinkDeviceUnbindAPIRequest(v *AlibabaAlinkDeviceUnbindAPIRequest) {
	v.Reset()
	poolAlibabaAlinkDeviceUnbindAPIRequest.Put(v)
}
