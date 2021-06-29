package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询设备分组 APIRequest
alibaba.campus.device.openapi.getdevicelist

多条件查询设备分组
*/
type AlibabaCampusDeviceOpenapiGetdevicelistRequest struct {
    model.Params

    // 请求发送端信息
    workBenchContext   *WorkBenchContext 

    // 多条件查询对象
    query   *DeviceApiQuery 

}

func NewAlibabaCampusDeviceOpenapiGetdevicelistRequest() *AlibabaCampusDeviceOpenapiGetdevicelistRequest{
    return &AlibabaCampusDeviceOpenapiGetdevicelistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGetdevicelistRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getdevicelist"
}

func (r AlibabaCampusDeviceOpenapiGetdevicelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGetdevicelistRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicelistRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGetdevicelistRequest) SetQuery(query *DeviceApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicelistRequest) GetQuery() *DeviceApiQuery {
    return r.query
}

