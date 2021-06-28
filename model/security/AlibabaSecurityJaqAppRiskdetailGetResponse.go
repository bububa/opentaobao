package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
应用风险详细信息查询接口 APIResponse
alibaba.security.jaq.app.riskdetail.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
type AlibabaSecurityJaqAppRiskdetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqAppRiskdetailGetResponse `json:"alibaba_security_jaq_app_riskdetail_get_response,omitempty"` 
    AlibabaSecurityJaqAppRiskdetailGetResponse
}

/* model for simplify = false
type AlibabaSecurityJaqAppRiskdetailGetResponse struct {

    // 风险详情
    
    Result  *struct {
        RiskDetail  *RiskDetail `json:"risk_detail,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSecurityJaqAppRiskdetailGetResponse struct {

    // 风险详情
    Result   *RiskDetail `json:"result,omitempty"`

}
