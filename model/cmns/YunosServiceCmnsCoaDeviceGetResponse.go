package cmns

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备详情查询 APIResponse
yunos.service.cmns.coa.device.get

第三方应用开发者调用此接口查询设备详情
*/
type YunosServiceCmnsCoaDeviceGetAPIResponse struct {
    model.CommonResponse
    // Response *YunosServiceCmnsCoaDeviceGetResponse `json:"yunos_service_cmns_coa_device_get_response,omitempty"` 
    YunosServiceCmnsCoaDeviceGetResponse
}

/* model for simplify = false
type YunosServiceCmnsCoaDeviceGetResponse struct {

    // 设备详情
    
    DeviceList  struct {
        DeviceResult  []DeviceResult `json:"device_result,omitempty"`
    } `json:"device_list,omitempty"`
    

    // 接口查询出错提示信息
    
    Message   string `json:"message,omitempty"`
    

    // 200表示查询成功
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

type YunosServiceCmnsCoaDeviceGetResponse struct {

    // 设备详情
    DeviceList   []DeviceResult `json:"device_list,omitempty"`

    // 接口查询出错提示信息
    Message   string `json:"message,omitempty"`

    // 200表示查询成功
    Status   int64 `json:"status,omitempty"`

}
