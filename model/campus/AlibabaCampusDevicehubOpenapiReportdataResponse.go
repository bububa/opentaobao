package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备数据上报 APIResponse
alibaba.campus.devicehub.openapi.reportdata

设备数据上报
*/
type AlibabaCampusDevicehubOpenapiReportdataAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDevicehubOpenapiReportdataResponse
}

type AlibabaCampusDevicehubOpenapiReportdataResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_devicehub_openapi_reportdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 自动生成
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
