package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest
获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
alibaba.campus.device.openapi.getsimpledevice

获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) */
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备uuid
	_uuid string
}

// NewAlibabaCampusDeviceOpenapiGetsimpledeviceRequest 初始化AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetsimpledeviceRequest() *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getsimpledevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// Get WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// Set is Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest) GetUuid() string {
	return r._uuid
}
