package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
应用风险详细信息批量查询接口 APIResponse
alibaba.security.jaq.app.riskdetailbatch.get

用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
type AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqAppRiskdetailbatchGetResponse `json:"alibaba_security_jaq_app_riskdetailbatch_get_response,omitempty"` 
    AlibabaSecurityJaqAppRiskdetailbatchGetResponse
}

/* model for simplify = false
type AlibabaSecurityJaqAppRiskdetailbatchGetResponse struct {

    // 批量扫描风险详情
    
    Result  *struct {
        RiskDetailBatch  *RiskDetailBatch `json:"risk_detail_batch,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSecurityJaqAppRiskdetailbatchGetResponse struct {

    // 批量扫描风险详情
    Result   *RiskDetailBatch `json:"result,omitempty"`

}
