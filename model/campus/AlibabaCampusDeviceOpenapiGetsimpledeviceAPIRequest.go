package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigetsimpledeviceAPIRequest 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
// alibaba.campus.device.openapi.getsimpledevice
//
// 获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
type AlibabacampusdeviceopenapigetsimpledeviceAPIRequest struct {
	model.Params
	// 设备uuid
	_uuid string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabacampusdeviceopenapigetsimpledeviceRequest 初始化AlibabacampusdeviceopenapigetsimpledeviceAPIRequest对象
func NewAlibabacampusdeviceopenapigetsimpledeviceRequest() *AlibabacampusdeviceopenapigetsimpledeviceAPIRequest {
	return &AlibabacampusdeviceopenapigetsimpledeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getsimpledevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) GetUuid() string {
	return r._uuid
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigetsimpledeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}
