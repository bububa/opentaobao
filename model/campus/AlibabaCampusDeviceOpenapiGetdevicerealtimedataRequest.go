package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定设备下指定参数的实时值 APIRequest
alibaba.campus.device.openapi.getdevicerealtimedata

获取指定设备下指定参数的实时值
*/
type AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest struct {
    model.Params

    // 请求端信息
    workBenchContext   *WorkBenchContext 

    // 设备uuid
    uuid   string 

    // 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
    propertyCode   string 

}

func NewAlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest{
    return &AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getdevicerealtimedata"
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) SetPropertyCode(propertyCode string) error {
    r.propertyCode = propertyCode
    r.Set("property_code", propertyCode)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetPropertyCode() string {
    return r.propertyCode
}

