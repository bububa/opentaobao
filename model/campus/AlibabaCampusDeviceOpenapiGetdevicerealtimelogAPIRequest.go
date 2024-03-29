package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest 根据设备uuid获取设备采集信息 API请求
// alibaba.campus.device.openapi.getdevicerealtimelog
//
// 根据设备uuid获取设备采集信息
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest struct {
	model.Params
	// 设备uuid
	_uuid string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest 初始化AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) Reset() {
	r._uuid = ""
	r._workBenchContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicerealtimelog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetUuid() string {
	return r._uuid
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

var poolAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest
func GetAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest {
	return poolAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest.Get().(*AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest 将 AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest(v *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest.Put(v)
}
