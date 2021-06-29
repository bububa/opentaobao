package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备授权 APIResponse
alibaba.ailabs.tmallgenie.auth.gettoken

获取设备授权码
*/
type AlibabaAilabsTmallgenieAuthGettokenAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthGettokenResponse
}

type AlibabaAilabsTmallgenieAuthGettokenResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_gettoken_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设备注册结果
    
    RegisterResult   *RegisterInfoVo `json:"register_result,omitempty" xml:"register_result,omitempty"`

    
    // 授权结果
    
    AuthResult   *DeviceTokenVo `json:"auth_result,omitempty" xml:"auth_result,omitempty"`

    
}
