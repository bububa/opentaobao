package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid操作设备 API请求
alibaba.campus.device.openapi.operatedevice

根据uuid操作设备
*/
type AlibabaCampusDeviceOpenapiOperatedeviceRequest struct {
    model.Params
    // 请求发送端信息
    _workBenchContext   *WorkBenchContext
    // 设备序列号uuid
    _uuid   string
    // 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
    _propertyCode   string
    // 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
    _value   string
}

// 初始化AlibabaCampusDeviceOpenapiOperatedeviceRequest对象
func NewAlibabaCampusDeviceOpenapiOperatedeviceRequest() *AlibabaCampusDeviceOpenapiOperatedeviceRequest{
    return &AlibabaCampusDeviceOpenapiOperatedeviceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.operatedevice"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
// Uuid Setter
// 设备序列号uuid
func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetUuid() string {
    return r._uuid
}
// PropertyCode Setter
// 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetPropertyCode(_propertyCode string) error {
    r._propertyCode = _propertyCode
    r.Set("property_code", _propertyCode)
    return nil
}

// PropertyCode Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetPropertyCode() string {
    return r._propertyCode
}
// Value Setter
// 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
func (r *AlibabaCampusDeviceOpenapiOperatedeviceRequest) SetValue(_value string) error {
    r._value = _value
    r.Set("value", _value)
    return nil
}

// Value Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceRequest) GetValue() string {
    return r._value
}
