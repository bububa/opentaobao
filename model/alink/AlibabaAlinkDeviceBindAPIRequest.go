package alink

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceBindAPIRequest 绑定设备 API请求
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
type AlibabaAlinkDeviceBindAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaAlinkDeviceBindRequest 初始化AlibabaAlinkDeviceBindAPIRequest对象
func NewAlibabaAlinkDeviceBindRequest() *AlibabaAlinkDeviceBindAPIRequest {
	return &AlibabaAlinkDeviceBindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlinkDeviceBindAPIRequest) Reset() {
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkDeviceBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceBindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceBindAPIRequest) GetUuid() string {
	return r._uuid
}

var poolAlibabaAlinkDeviceBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlinkDeviceBindRequest()
	},
}

// GetAlibabaAlinkDeviceBindRequest 从 sync.Pool 获取 AlibabaAlinkDeviceBindAPIRequest
func GetAlibabaAlinkDeviceBindAPIRequest() *AlibabaAlinkDeviceBindAPIRequest {
	return poolAlibabaAlinkDeviceBindAPIRequest.Get().(*AlibabaAlinkDeviceBindAPIRequest)
}

// ReleaseAlibabaAlinkDeviceBindAPIRequest 将 AlibabaAlinkDeviceBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlinkDeviceBindAPIRequest(v *AlibabaAlinkDeviceBindAPIRequest) {
	v.Reset()
	poolAlibabaAlinkDeviceBindAPIRequest.Put(v)
}
