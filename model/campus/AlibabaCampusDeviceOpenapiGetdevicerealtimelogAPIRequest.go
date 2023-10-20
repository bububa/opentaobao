package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest 根据设备uuid获取设备采集信息 API请求
// alibaba.campus.device.openapi.getdevicerealtimelog
//
// 根据设备uuid获取设备采集信息
type AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest struct {
	model.Params
	// 设备uuid
	_uuid string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabacampusdeviceopenapigetdevicerealtimelogRequest 初始化AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest对象
func NewAlibabacampusdeviceopenapigetdevicerealtimelogRequest() *AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest {
	return &AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicerealtimelog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) GetUuid() string {
	return r._uuid
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}
