package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceInvoiceapplyGetAPIResponse 获取商家的开票申请 API返回值
// alibaba.einvoice.invoiceapply.get
//
// 开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
type AlibabaEinvoiceInvoiceapplyGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceInvoiceapplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceInvoiceapplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceInvoiceapplyGetAPIResponseModel).Reset()
}

// AlibabaEinvoiceInvoiceapplyGetAPIResponseModel is 获取商家的开票申请 成功返回结果
type AlibabaEinvoiceInvoiceapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_invoiceapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票申请查询结果
	Result *UserInvoiceApplyDto `json:"result,omitempty" xml:"result,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceInvoiceapplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
	m.TotalCount = 0
}

var poolAlibabaEinvoiceInvoiceapplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceInvoiceapplyGetAPIResponse)
	},
}

// GetAlibabaEinvoiceInvoiceapplyGetAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceInvoiceapplyGetAPIResponse
func GetAlibabaEinvoiceInvoiceapplyGetAPIResponse() *AlibabaEinvoiceInvoiceapplyGetAPIResponse {
	return poolAlibabaEinvoiceInvoiceapplyGetAPIResponse.Get().(*AlibabaEinvoiceInvoiceapplyGetAPIResponse)
}

// ReleaseAlibabaEinvoiceInvoiceapplyGetAPIResponse 将 AlibabaEinvoiceInvoiceapplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceInvoiceapplyGetAPIResponse(v *AlibabaEinvoiceInvoiceapplyGetAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceInvoiceapplyGetAPIResponse.Put(v)
}
