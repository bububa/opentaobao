package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备历史数据 API返回值 
alibaba.campus.device.openapi.gethistorydata

查询历史数据的接口
*/
type AlibabaCampusDeviceOpenapiGethistorydataAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiGethistorydataResponse
}

// 查询设备历史数据 成功返回结果
type AlibabaCampusDeviceOpenapiGethistorydataResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_gethistorydata_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
