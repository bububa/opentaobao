package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备历史数据 APIRequest
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

func NewAlibabaCampusDeviceOpenapiGethistorydataRequest() *AlibabaCampusDeviceOpenapiGethistorydataRequest{
    return &AlibabaCampusDeviceOpenapiGethistorydataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.gethistorydata"
}

func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGethistorydataRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGethistorydataRequest) SetQuery(query *DeviceDataApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGethistorydataRequest) GetQuery() *DeviceDataApiQuery {
    return r.query
}

