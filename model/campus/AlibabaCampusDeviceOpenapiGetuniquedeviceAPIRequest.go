package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigetuniquedeviceAPIRequest 根据设备uuid获取设备信息 API请求
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
type AlibabacampusdeviceopenapigetuniquedeviceAPIRequest struct {
	model.Params
	// 设备序列号uuid
	_uuid string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabacampusdeviceopenapigetuniquedeviceRequest 初始化AlibabacampusdeviceopenapigetuniquedeviceAPIRequest对象
func NewAlibabacampusdeviceopenapigetuniquedeviceRequest() *AlibabacampusdeviceopenapigetuniquedeviceAPIRequest {
	return &AlibabacampusdeviceopenapigetuniquedeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getuniquedevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备序列号uuid
func (r *AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) GetUuid() string {
	return r._uuid
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigetuniquedeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}
