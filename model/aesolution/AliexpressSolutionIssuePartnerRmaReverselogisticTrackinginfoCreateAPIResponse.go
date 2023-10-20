package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create API返回值
// aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create
//
// Receives information about reverse logistics tracking info
type AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponseModel).Reset()
}

// AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponseModel is aliexpress.solution.issue.partner.rma.reverselogistic.trackinginfo.create 成功返回结果
type AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_reverselogistic_trackinginfo_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Error code. 0 value is no error.
	CodeError string `json:"code_error,omitempty" xml:"code_error,omitempty"`
	// Error description
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CodeError = ""
	m.ErrorDescription = ""
}

var poolAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse)
	},
}

// GetAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse
func GetAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse() *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse {
	return poolAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse.Get().(*AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse)
}

// ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse 将 AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse(v *AliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaReverselogisticTrackinginfoCreateAPIResponse.Put(v)
}
