package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备详情查询 APIResponse
yunos.service.cmns.coa.device.get

第三方应用开发者调用此接口查询设备详情
*/
type YunosServiceCmnsCoaDeviceGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_device_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 设备详情
    
    DeviceList   []DeviceResult `json:"device_list,omitempty" xml:"