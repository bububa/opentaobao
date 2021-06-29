package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户设备列表 APIResponse
alibaba.ailabs.tmallgenie.auth.device.list

通过此接口获取用户绑定的设备信息列表
*/
type AlibabaAilabsTmallgenieAuthDeviceListAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthDeviceListResponse
}

type AlibabaAilabsTmallgenieAuthDeviceListResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设备列表
    
    Devices   []AlibabaAilabsTmallgenieAuthDeviceListResult `json:"devices,omitempty" xml:"devices>alibaba_ailabs_tmallgenie_auth_device_list_result,omitempty"`
    
    
}
