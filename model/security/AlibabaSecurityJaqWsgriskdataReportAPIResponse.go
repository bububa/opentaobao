package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqwsgriskdatareportAPIResponse 无线保镖SDK风控数据上报 API返回值
// alibaba.security.jaq.wsgriskdata.report
//
// 无线保镖sdk根据用户的需要，上报数据到聚安全云端
type AlibabasecurityjaqwsgriskdatareportAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqwsgriskdatareportAPIResponseModel
}

// AlibabasecurityjaqwsgriskdatareportAPIResponseModel is 无线保镖SDK风控数据上报 成功返回结果
type AlibabasecurityjaqwsgriskdatareportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_wsgriskdata_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无线保镖sdk上报数据的返回结果
	JaqWsgRiskReportResult *JaqWsgReportResult `json:"jaq_wsg_risk_report_result,omitempty" xml:"jaq_wsg_risk_report_result,omitempty"`
}
