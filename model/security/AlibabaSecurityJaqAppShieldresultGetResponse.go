package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户查询加固结果 APIResponse
alibaba.security.jaq.app.shieldresult.get

用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
*/
type AlibabaSecurityJaqAppShieldresultGetAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqAppShieldresultGetResponse
}

type AlibabaSecurityJaqAppShieldresultGetResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_app_shieldresult_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 应用加固结果
    
    Result   *ShieldResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
