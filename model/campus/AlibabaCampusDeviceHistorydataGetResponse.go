package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备历史数据批量获取 APIResponse
alibaba.campus.device.historydata.get

设备历史数据批量获取
*/
type AlibabaCampusDeviceHistorydataGetAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceHistorydataGetResponse
}

type AlibabaCampusDeviceHistorydataGetResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_historydata_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
