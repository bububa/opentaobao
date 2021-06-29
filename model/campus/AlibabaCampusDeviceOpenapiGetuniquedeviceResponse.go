package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备uuid获取设备信息 APIResponse
alibaba.campus.device.openapi.getuniquedevice

根据设备uuid获取设备信息
*/
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGetuniquedeviceResponse
}

type AlibabaCampusDeviceOpenapiGetuniquedeviceResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_getuniquedevice_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
