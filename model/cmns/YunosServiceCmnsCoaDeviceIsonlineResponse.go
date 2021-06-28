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
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_device_isonline_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   int64 `json:"data,omitempty" xml:"