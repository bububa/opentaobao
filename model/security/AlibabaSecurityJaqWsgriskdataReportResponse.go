package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
无线保镖SDK风控数据上报 APIResponse
alibaba.security.jaq.wsgriskdata.report

无线保镖sdk根据用户的需要，上报数据到聚安全云端
*/
type AlibabaSecurityJaqWsgriskdataReportAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqWsgriskdataReportResponse `json:"alibaba_security_jaq_wsgriskdata_report_response,omitempty"`
}

type AlibabaSecurityJaqWsgriskdataReportResponse struct {

    // 无线保镖sdk上报数据的返回结果
    JaqWsgRiskReportResult   *JaqWsgReportResult `json:"jaq_wsg_risk_report_result,omitempty"`

}
