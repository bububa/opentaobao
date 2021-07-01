package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API返回值 
alibaba.campus.device.openapi.getsimpledevice

获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
*/
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel
}

// 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) 成功返回结果
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_getsimpledevice_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
