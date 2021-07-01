package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest
根据设备uuid获取设备采集信息 API请求
alibaba.campus.device.openapi.getdevicerealtimelog

根据设备uuid获取设备采集信息 */
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备uuid
	_uuid string
}

// NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest 初始化AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicerealtimelog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// Get WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// Set is Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetUuid() string {
	return r._uuid
}
