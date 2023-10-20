package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse 触发odps任务离线查询公益宝贝开票对账明细 API返回值
// alibaba.csr.donate.invoice.querytoblockchainoss
//
// 提供给蚂蚁链上公益团队，用于触发odps任务离线查询公益宝贝开票对账明细
type AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse struct {
	model.CommonResponse
	AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponseModel).Reset()
}

// AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponseModel is 触发odps任务离线查询公益宝贝开票对账明细 成功返回结果
type AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_donate_invoice_querytoblockchainoss_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 错误码
	FailCode int64 `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	HasSuccess bool `json:"has_success,omitempty" xml:"has_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailMsg = ""
	m.FailCode = 0
	m.HasSuccess = false
}

var poolAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse)
	},
}

// GetAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse 从 sync.Pool 获取 AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse
func GetAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse() *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse {
	return poolAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse.Get().(*AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse)
}

// ReleaseAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse 将 AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse(v *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse) {
	v.Reset()
	poolAlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse.Put(v)
}
