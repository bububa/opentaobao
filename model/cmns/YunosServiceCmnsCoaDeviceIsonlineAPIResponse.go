package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id查询设备是否在线 API返回值 
yunos.service.cmns.coa.device.isonline

根据设备id查询设备是否在线
*/
type YunosServiceCmnsCoaDeviceIsonlineAPIResponse struct {
    model.CommonResponse
    YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel
}

// 根据设备id查询设备是否在线 成功返回结果
type YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_service_cmns_coa_device_isonline_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`
    // msg
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // status
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
