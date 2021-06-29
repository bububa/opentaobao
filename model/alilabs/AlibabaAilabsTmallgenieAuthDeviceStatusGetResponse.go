package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备状态查询 APIResponse
alibaba.ailabs.tmallgenie.auth.device.status.get

提供给天猫精灵定制机厂商 查询设备在线状态值
*/
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthDeviceStatusGetResponse
}

type AlibabaAilabsTmallgenieAuthDeviceStatusGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_status_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口对象封装
    
    Result   *AlibabaAilabsTmallgenieAuthDeviceStatusGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
