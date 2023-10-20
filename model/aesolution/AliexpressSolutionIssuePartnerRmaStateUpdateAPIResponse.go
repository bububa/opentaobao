package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse aliexpress.solution.issue.partner.rma.state.update API返回值
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
type AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponseModel).Reset()
}

// AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponseModel is aliexpress.solution.issue.partner.rma.state.update 成功返回结果
type AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_state_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error code. 0 means no error
	CodeError string `json:"code_error,omitempty" xml:"code_error,omitempty"`
	// error description
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CodeError = ""
	m.ErrorDescription = ""
}

var poolAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse)
	},
}

// GetAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse
func GetAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse() *AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse {
	return poolAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse.Get().(*AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse)
}

// ReleaseAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse 将 AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse(v *AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse.Put(v)
}
