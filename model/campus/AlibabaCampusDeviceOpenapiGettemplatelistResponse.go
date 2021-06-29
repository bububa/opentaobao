package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备模板 APIResponse
alibaba.campus.device.openapi.gettemplatelist

查询设备模板信息
*/
type AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGettemplatelistResponse
}

type AlibabaCampusDeviceOpenapiGettemplatelistResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_gettemplatelist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
