package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) APIResponse
alibaba.campus.device.openapi.getsimpledevicelist

查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
*/
type AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGetsimpledevicelistResponse
}

type AlibabaCampusDeviceOpenapiGetsimpledevicelistResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_getsimpledevicelist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
