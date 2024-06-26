package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
// alibaba.campus.device.openapi.getsimpledevice
//
// 获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest struct {
	model.Params
	// 设备uuid
	_uuid string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabaCampusDeviceOpenapiGetsimpledeviceRequest 初始化AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetsimpledeviceRequest() *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) Reset() {
	r._uuid = ""
	r._workBenchContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getsimpledevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetUuid() string {
	return r._uuid
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

var poolAlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiGetsimpledeviceRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiGetsimpledeviceRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest
func GetAlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest() *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest {
	return poolAlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest.Get().(*AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest 将 AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest(v *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest.Put(v)
}
