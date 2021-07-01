package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下发设备的分页接口(无需AOP控制) API请求
alibaba.campus.device.getdeviceforquery

下发设备的分页接口(发布在TOP上，connect调用，无需AOP控制)
*/
type AlibabaCampusDeviceGetdeviceforqueryAPIRequest struct {
    model.Params
    // 系统自动生成
    _query   *DeviceApiQuery
    // 平台统一参数
    _workBenchContext   *WorkBenchContext
}

// 初始化AlibabaCampusDeviceGetdeviceforqueryAPIRequest对象
func NewAlibabaCampusDeviceGetdeviceforqueryRequest() *AlibabaCampusDeviceGetdeviceforqueryAPIRequest{
    return &AlibabaCampusDeviceGetdeviceforqueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceGetdeviceforqueryAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.device.getdeviceforquery"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceGetdeviceforqueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 系统自动生成
func (r *AlibabaCampusDeviceGetdeviceforqueryAPIRequest) SetQuery(_query *DeviceApiQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaCampusDeviceGetdeviceforqueryAPIRequest) GetQuery() *DeviceApiQuery {
    return r._query
}
// WorkBenchContext Setter
// 平台统一参数
func (r *AlibabaCampusDeviceGetdeviceforqueryAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceGetdeviceforqueryAPIRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
