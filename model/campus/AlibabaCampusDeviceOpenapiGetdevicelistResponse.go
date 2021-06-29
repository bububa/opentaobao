package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询设备分组 APIResponse
alibaba.campus.device.openapi.getdevicelist

多条件查询设备分组
*/
type AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGetdevicelistResponse
}

type AlibabaCampusDeviceOpenapiGetdevicelistResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_getdevicelist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
