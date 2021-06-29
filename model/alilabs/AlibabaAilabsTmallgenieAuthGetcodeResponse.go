package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取token APIResponse
alibaba.ailabs.tmallgenie.auth.getcode

获取天猫精灵authCode
*/
type AlibabaAilabsTmallgenieAuthGetcodeAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthGetcodeResponse
}

type AlibabaAilabsTmallgenieAuthGetcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_getcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 授权码 json 串，包含授权码和二维码URL
    
    AuthCode   string `json:"auth_code,omitempty" xml:"auth_code,omitempty"`

    
}
