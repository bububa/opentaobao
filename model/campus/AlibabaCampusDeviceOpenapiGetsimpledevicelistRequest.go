package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) APIRequest
alibaba.campus.device.openapi.getsimpledevicelist

查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
*/
type AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest struct {
    model.Params

    // 请求发送端信息
    workBenchContext   *WorkBenchContext 

    // 多条件查询对象
    query   *DeviceApiQuery 

}

func NewAlibabaCampusDeviceOpenapiGetsimpledevicelistRequest() *AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest{
    return &AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getsimpledevicelist"
}

func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) SetQuery(query *DeviceApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistRequest) GetQuery() *DeviceApiQuery {
    return r.query
}

