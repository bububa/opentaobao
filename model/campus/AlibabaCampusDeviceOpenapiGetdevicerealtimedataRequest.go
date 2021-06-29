package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定设备下指定参数的实时值 API请求
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

// 初始化AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest{
    return &AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getdevicerealtimedata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetUuid() string {
    return r.uuid
}
// PropertyCode Setter
// 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) SetPropertyCode(propertyCode string) error {
    r.propertyCode = propertyCode
    r.Set("property_code", propertyCode)
    return nil
}

// PropertyCode Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest) GetPropertyCode() string {
    return r.propertyCode
}
