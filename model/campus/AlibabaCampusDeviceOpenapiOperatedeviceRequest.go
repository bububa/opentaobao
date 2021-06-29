package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid操作设备 APIRequest
alibaba.campus.device.openapi.operatedevice

根据uuid操作设备
*/
type AlibabaCampusDeviceOpenapiOperatedeviceRequest struct {
    model.Params

    // 请求发送端信息
    workBenchContext   *WorkBenchContext 

    // 设备序列号uuid
    uuid   string 

    // 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
    propertyCode   string 

    // 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
    value   string 

}

func NewAlibabaCampusDeviceOpenapiOperatedeviceRequest() *AlibabaCampusDeviceOpenapiOperatedeviceRequest{
    return &AlibabaCampusDeviceOpenapiOperatedeviceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.operatedevice"
}

func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetPropertyCode(propertyCode string) error {
    r.propertyCode = propertyCode
    r.Set("property_code", propertyCode)
    return nil
}

func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetPropertyCode() string {
    return r.propertyCode
}

func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetValue(value string) error {
    r.value = value
    r.Set("value", value)
    return nil
}

func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetValue() string {
    return r.value
}

