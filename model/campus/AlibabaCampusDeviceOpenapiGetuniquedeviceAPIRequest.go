package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest 根据设备uuid获取设备信息 API请求
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备序列号uuid
	_uuid string
}

// NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest 初始化AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest() *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getuniquedevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// Get WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// Set is Uuid Setter
// 设备序列号uuid
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetUuid() string {
	return r._uuid
}
