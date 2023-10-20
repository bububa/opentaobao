package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse aliexpress.solution.issue.partner.rma.reverselogistic.state.update API返回值
// aliexpress.solution.issue.partner.rma.reverselogistic.state.update
//
// Updates the reverse logistics state for after sales services
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponseModel).Reset()
}

// AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponseModel is aliexpress.solution.issue.partner.rma.reverselogistic.state.update 成功返回结果
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_reverselogistic_state_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Error code. 0 value is no error.
	CodeError string `json:"code_error,omitempty" xml:"code_error,omitempty"`
	// Error description
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CodeError = ""
	m.ErrorDescription = ""
}

var poolAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse)
	},
}

// GetAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse
func GetAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse() *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse {
	return poolAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse.Get().(*AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse)
}

// ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse 将 AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse(v *AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse.Put(v)
}
