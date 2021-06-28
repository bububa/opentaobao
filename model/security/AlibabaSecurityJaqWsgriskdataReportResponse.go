package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无线保镖SDK风控数据上报 APIResponse
alibaba.security.jaq.wsgriskdata.report

无线保镖sdk根据用户的需要，上报数据到聚安全云端
*/
type AlibabaSecurityJaqWsgriskdataReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_security_jaq_wsgriskdata_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 无线保镖sdk上报数据的返回结果
    
    JaqWsgRiskReportResult   *JaqWsgReportResult `json:"jaq_wsg_risk_report_result,omitempty" xml:"