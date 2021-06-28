package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
应用加固接口 APIResponse
alibaba.security.jaq.app.shield

提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
*/
type AlibabaSecurityJaqAppShieldAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqAppShieldResponse
}

type AlibabaSecurityJaqAppShieldResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_app_shield_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 加固任务信息
    
    Result   *ScanTaskInfo `json:"result,omitempty" xml:"result,omitempty"`

    
}
