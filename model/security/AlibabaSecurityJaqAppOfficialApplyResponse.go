package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全官方应用申请 APIResponse
alibaba.security.jaq.app.official.apply

官方应用申请接口
*/
type AlibabaSecurityJaqAppOfficialApplyAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqAppOfficialApplyResponse
}

type AlibabaSecurityJaqAppOfficialApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_app_official_apply_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 申请结果
    
    Result   *OfficialAppApplyResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
