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
type AlibabaCampusDeviceGetdeviceforqueryRequest struct {
    model.Params
    // 系统自动生成
    _query   *DeviceApiQuery
    // 平台统一参数
    _workBenchContext   *WorkBenchContext
}

// 初始化AlibabaCampusDeviceGetdeviceforqueryRequest对象
func NewAlibabaCampusDeviceGetdeviceforqueryRequest() *AlibabaCampusDeviceGetdeviceforqueryRequest{
    return &AlibabaCampusDeviceGetdeviceforqueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceGetdeviceforqueryRequest) GetApiMethodName() string {
    return "alibaba.campus.device.getdeviceforquery"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceGetdeviceforqueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 系统自动生成
func (r *AlibabaCampusDeviceGetdeviceforqueryRequest) SetQuery(_query *DeviceApiQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaCampusDeviceGetdeviceforqueryRequest) GetQuery() *DeviceApiQuery {
    return r._query
}
// WorkBenchContext Setter
// 平台统一参数
func (r *AlibabaCampusDeviceGetdeviceforqueryRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceGetdeviceforqueryRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
