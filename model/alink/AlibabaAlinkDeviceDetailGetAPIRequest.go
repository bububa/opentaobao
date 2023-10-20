package alink

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceDetailGetAPIRequest 获取设备详情 API请求
// alibaba.alink.device.detail.get
//
// 阿里智能获取设备详情
type AlibabaAlinkDeviceDetailGetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaAlinkDeviceDetailGetRequest 初始化AlibabaAlinkDeviceDetailGetAPIRequest对象
func NewAlibabaAlinkDeviceDetailGetRequest() *AlibabaAlinkDeviceDetailGetAPIRequest {
	return &AlibabaAlinkDeviceDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlinkDeviceDetailGetAPIRequest) Reset() {
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceDetailGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceDetailGetAPIRequest) GetUuid() string {
	return r._uuid
}

var poolAlibabaAlinkDeviceDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlinkDeviceDetailGetRequest()
	},
}

// GetAlibabaAlinkDeviceDetailGetRequest 从 sync.Pool 获取 AlibabaAlinkDeviceDetailGetAPIRequest
func GetAlibabaAlinkDeviceDetailGetAPIRequest() *AlibabaAlinkDeviceDetailGetAPIRequest {
	return poolAlibabaAlinkDeviceDetailGetAPIRequest.Get().(*AlibabaAlinkDeviceDetailGetAPIRequest)
}

// ReleaseAlibabaAlinkDeviceDetailGetAPIRequest 将 AlibabaAlinkDeviceDetailGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlinkDeviceDetailGetAPIRequest(v *AlibabaAlinkDeviceDetailGetAPIRequest) {
	v.Reset()
	poolAlibabaAlinkDeviceDetailGetAPIRequest.Put(v)
}
