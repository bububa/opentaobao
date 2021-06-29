package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险扫描批量提交接口 APIResponse
alibaba.security.jaq.app.risk.scanbatch

批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果
*/
type AlibabaSecurityJaqAppRiskScanbatchAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqAppRiskScanbatchResponse
}

type AlibabaSecurityJaqAppRiskScanbatchResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_app_risk_scanbatch_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 扫描任务信息
    
    Result   *TaskInfo `json:"result,omitempty" xml:"result,omitempty"`

    
}
