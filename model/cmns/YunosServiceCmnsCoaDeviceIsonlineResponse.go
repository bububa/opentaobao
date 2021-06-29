package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id查询设备是否在线 APIResponse
yunos.service.cmns.coa.device.isonline

根据设备id查询设备是否在线
*/
type YunosServiceCmnsCoaDeviceIsonlineAPIResponse struct {
    model.CommonResponse
    YunosServiceCmnsCoaDeviceIsonlineResponse
}

type YunosServiceCmnsCoaDeviceIsonlineResponse struct {
    XMLName xml.Name `xml:"yunos_service_cmns_coa_device_isonline_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`

    
    // msg
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // status
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`

    
}
