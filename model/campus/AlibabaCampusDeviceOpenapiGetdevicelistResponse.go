package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多条件查询设备分组 API返回值 
alibaba.campus.device.openapi.getdevicelist

多条件查询设备分组
*/
type AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGetdevicelistResponse
}

// 多条件查询设备分组 成功返回结果
type AlibabaCampusDeviceOpenapiGetdevicelistResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_getdevicelist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
