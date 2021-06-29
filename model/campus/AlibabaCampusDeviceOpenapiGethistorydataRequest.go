package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备历史数据 API请求
alibaba.campus.device.openapi.gethistorydata

查询历史数据的接口
*/
type AlibabaCampusDeviceOpenapiGethistorydataRequest struct {
    model.Params
    // 请求端信息
    workBenchContext   *WorkBenchContext
    // 历史数据查询对象
    query   *DeviceDataApiQuery
}

// 初始化AlibabaCampusDeviceOpenapiGethistorydataRequest对象
func NewAlibabaCampusDeviceOpenapiGethistorydataRequest() *AlibabaCampusDeviceOpenapiGethistorydataRequest{
    return &AlibabaCampusDeviceOpenapiGethistorydataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.gethistorydata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求端信息
func (r *AlibabaCampusDeviceOpenapiGethistorydataRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// Query Setter
// 历史数据查询对象
func (r *AlibabaCampusDeviceOpenapiGethistorydataRequest) SetQuery(query *DeviceDataApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetQuery() *DeviceDataApiQuery {
    return r.query
}
