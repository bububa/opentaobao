package cmns

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id查询设备是否在线 APIResponse
yunos.service.cmns.coa.device.isonline

根据设备id查询设备是否在线
*/
type YunosServiceCmnsCoaDeviceIsonlineAPIResponse struct {
    model.CommonResponse
    // Response *YunosServiceCmnsCoaDeviceIsonlineResponse `json:"yunos_service_cmns_coa_device_isonline_response,omitempty"` 
    YunosServiceCmnsCoaDeviceIsonlineResponse
}

/* model for simplify = false
type YunosServiceCmnsCoaDeviceIsonlineResponse struct {

    // data
    
    Data   int64 `json:"data,omitempty"`
    

    // msg
    
    Message   string `json:"message,omitempty"`
    

    // status
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

type YunosServiceCmnsCoaDeviceIsonlineResponse struct {

    // data
    Data   int64 `json:"data,omitempty"`

    // msg
    Message   string `json:"message,omitempty"`

    // status
    Status   int64 `json:"status,omitempty"`

}
