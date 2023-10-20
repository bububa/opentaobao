package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest 根据设备uuid获取设备信息 API请求
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest struct {
	model.Params
	// 设备序列号uuid
	_uuid string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest 初始化AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest() *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) Reset() {
	r._uuid = ""
	r._workBenchContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getuniquedevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备序列号uuid
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetUuid() string {
	return r._uuid
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

var poolAlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiGetuniquedeviceRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest
func GetAlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest() *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest {
	return poolAlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest.Get().(*AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest 将 AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest(v *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest.Put(v)
}
