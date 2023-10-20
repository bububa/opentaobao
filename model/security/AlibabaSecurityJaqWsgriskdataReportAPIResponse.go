package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqWsgriskdataReportAPIResponse 无线保镖SDK风控数据上报 API返回值
// alibaba.security.jaq.wsgriskdata.report
//
// 无线保镖sdk根据用户的需要，上报数据到聚安全云端
type AlibabaSecurityJaqWsgriskdataReportAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqWsgriskdataReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqWsgriskdataReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqWsgriskdataReportAPIResponseModel).Reset()
}

// AlibabaSecurityJaqWsgriskdataReportAPIResponseModel is 无线保镖SDK风控数据上报 成功返回结果
type AlibabaSecurityJaqWsgriskdataReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_wsgriskdata_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无线保镖sdk上报数据的返回结果
	JaqWsgRiskReportResult *JaqWsgReportResult `json:"jaq_wsg_risk_report_result,omitempty" xml:"jaq_wsg_risk_report_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqWsgriskdataReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.JaqWsgRiskReportResult = nil
}

var poolAlibabaSecurityJaqWsgriskdataReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqWsgriskdataReportAPIResponse)
	},
}

// GetAlibabaSecurityJaqWsgriskdataReportAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqWsgriskdataReportAPIResponse
func GetAlibabaSecurityJaqWsgriskdataReportAPIResponse() *AlibabaSecurityJaqWsgriskdataReportAPIResponse {
	return poolAlibabaSecurityJaqWsgriskdataReportAPIResponse.Get().(*AlibabaSecurityJaqWsgriskdataReportAPIResponse)
}

// ReleaseAlibabaSecurityJaqWsgriskdataReportAPIResponse 将 AlibabaSecurityJaqWsgriskdataReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqWsgriskdataReportAPIResponse(v *AlibabaSecurityJaqWsgriskdataReportAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqWsgriskdataReportAPIResponse.Put(v)
}
