package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid操作设备 APIResponse
alibaba.campus.device.openapi.operatedevice

根据uuid操作设备
*/
type AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiOperatedeviceResponse
}

type AlibabaCampusDeviceOpenapiOperatedeviceResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_operatedevice_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
