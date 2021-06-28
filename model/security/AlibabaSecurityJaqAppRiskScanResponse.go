package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
应用风险扫描提交接口 APIResponse
alibaba.security.jaq.app.risk.scan

提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
*/
type AlibabaSecurityJaqAppRiskScanAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqAppRiskScanResponse `json:"alibaba_security_jaq_app_risk_scan_response,omitempty"` 
    AlibabaSecurityJaqAppRiskScanResponse
}

/* model for simplify = false
type AlibabaSecurityJaqAppRiskScanResponse struct {

    // 扫描任务信息
    
    Result  *struct {
        ScanTaskInfo  *ScanTaskInfo `json:"scan_task_info,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSecurityJaqAppRiskScanResponse struct {

    // 扫描任务信息
    Result   *ScanTaskInfo `json:"result,omitempty"`

}
