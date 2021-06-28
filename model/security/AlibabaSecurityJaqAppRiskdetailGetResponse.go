package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险详细信息查询接口 APIResponse
alibaba.security.jaq.app.riskdetail.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
type AlibabaSecurityJaqAppRiskdetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqAppRiskdetailGetResponse
}

type AlibabaSecurityJaqAppRiskdetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_app_riskdetail_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 风险详情
    
    Result   *RiskDetail `json:"result,omitempty" xml:"result,omitempty"`

    
}
