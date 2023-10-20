package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceInvoiceapplyUpdateAPIResponse 商家开票申请单状态回传 API返回值
// alibaba.einvoice.invoiceapply.update
//
// 开票服务商更新商家开票申请单状态
type AlibabaEinvoiceInvoiceapplyUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceInvoiceapplyUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceInvoiceapplyUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceInvoiceapplyUpdateAPIResponseModel).Reset()
}

// AlibabaEinvoiceInvoiceapplyUpdateAPIResponseModel is 商家开票申请单状态回传 成功返回结果
type AlibabaEinvoiceInvoiceapplyUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_invoiceapply_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 更新结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceInvoiceapplyUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TotalCount = 0
	m.Result = false
}

var poolAlibabaEinvoiceInvoiceapplyUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceInvoiceapplyUpdateAPIResponse)
	},
}

// GetAlibabaEinvoiceInvoiceapplyUpdateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceInvoiceapplyUpdateAPIResponse
func GetAlibabaEinvoiceInvoiceapplyUpdateAPIResponse() *AlibabaEinvoiceInvoiceapplyUpdateAPIResponse {
	return poolAlibabaEinvoiceInvoiceapplyUpdateAPIResponse.Get().(*AlibabaEinvoiceInvoiceapplyUpdateAPIResponse)
}

// ReleaseAlibabaEinvoiceInvoiceapplyUpdateAPIResponse 将 AlibabaEinvoiceInvoiceapplyUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceInvoiceapplyUpdateAPIResponse(v *AlibabaEinvoiceInvoiceapplyUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceInvoiceapplyUpdateAPIResponse.Put(v)
}
