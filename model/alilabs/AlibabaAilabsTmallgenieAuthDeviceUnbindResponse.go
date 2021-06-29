package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIResponse
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备
*/
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthDeviceUnbindResponse
}

type AlibabaAilabsTmallgenieAuthDeviceUnbindResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_unbind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *AlibabaAilabsTmallgenieAuthDeviceUnbindResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
