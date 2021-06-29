package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵淘宝登录授权绑定接口 APIResponse
alibaba.ailabs.tmallgenie.auth.taobaoauth

厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备
*/
type AlibabaAilabsTmallgenieAuthTaobaoauthAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthTaobaoauthResponse
}

type AlibabaAilabsTmallgenieAuthTaobaoauthResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_taobaoauth_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 注册结果
    
    RegisterResult   *RegisterInfoVo `json:"register_result,omitempty" xml:"register_result,omitempty"`

    
    // 授权结果
    
    AuthResult   *DeviceTokenVo `json:"auth_result,omitempty" xml:"auth_result,omitempty"`

    
}
