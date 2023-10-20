package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGethistorydataAPIRequest 查询设备历史数据 API请求
// alibaba.campus.device.openapi.gethistorydata
//
// 查询历史数据的接口
type AlibabaCampusDeviceOpenapiGethistorydataAPIRequest struct {
	model.Params
	// 请求端信息
	_workBenchContext *WorkBenchContext
	// 历史数据查询对象
	_query *DeviceDataApiQuery
}

// NewAlibabaCampusDeviceOpenapiGethistorydataRequest 初始化AlibabaCampusDeviceOpenapiGethistorydataAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGethistorydataRequest() *AlibabaCampusDeviceOpenapiGethistorydataAPIRequest {
	return &AlibabaCampusDeviceOpenapiGethistorydataAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) Reset() {
	r._workBenchContext = nil
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.gethistorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求端信息
func (r *AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 历史数据查询对象
func (r *AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) SetQuery(_query *DeviceDataApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) GetQuery() *DeviceDataApiQuery {
	return r._query
}

var poolAlibabaCampusDeviceOpenapiGethistorydataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiGethistorydataRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiGethistorydataRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGethistorydataAPIRequest
func GetAlibabaCampusDeviceOpenapiGethistorydataAPIRequest() *AlibabaCampusDeviceOpenapiGethistorydataAPIRequest {
	return poolAlibabaCampusDeviceOpenapiGethistorydataAPIRequest.Get().(*AlibabaCampusDeviceOpenapiGethistorydataAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiGethistorydataAPIRequest 将 AlibabaCampusDeviceOpenapiGethistorydataAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGethistorydataAPIRequest(v *AlibabaCampusDeviceOpenapiGethistorydataAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGethistorydataAPIRequest.Put(v)
}
