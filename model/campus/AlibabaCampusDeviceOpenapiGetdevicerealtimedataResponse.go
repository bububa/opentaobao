package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定设备下指定参数的实时值 APIResponse
alibaba.campus.device.openapi.getdevicerealtimedata

获取指定设备下指定参数的实时值
*/
type AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGetdevicerealtimedataResponse
}

type AlibabaCampusDeviceOpenapiGetdevicerealtimedataResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_getdevicerealtimedata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回查询结果
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
