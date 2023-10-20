package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse aliexpress.solution.issue.partner.rma.screening.create API返回值
// aliexpress.solution.issue.partner.rma.screening.create
//
// Receives information about screening results from after sales partners
type AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponseModel).Reset()
}

// AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponseModel is aliexpress.solution.issue.partner.rma.screening.create 成功返回结果
type AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_screening_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error code. 0 means no error
	CodeError string `json:"code_error,omitempty" xml:"code_error,omitempty"`
	// error description
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CodeError = ""
	m.ErrorDescription = ""
}

var poolAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse)
	},
}

// GetAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse
func GetAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse() *AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse {
	return poolAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse.Get().(*AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse)
}

// ReleaseAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse 将 AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse(v *AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse.Put(v)
}
