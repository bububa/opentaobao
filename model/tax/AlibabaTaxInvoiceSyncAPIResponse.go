package tax

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaxInvoiceSyncAPIResponse 第三方开票回调API API返回值
// alibaba.tax.invoice.sync
//
// 该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
type AlibabaTaxInvoiceSyncAPIResponse struct {
	model.CommonResponse
	AlibabaTaxInvoiceSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTaxInvoiceSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTaxInvoiceSyncAPIResponseModel).Reset()
}

// AlibabaTaxInvoiceSyncAPIResponseModel is 第三方开票回调API 成功返回结果
type AlibabaTaxInvoiceSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tax_invoice_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ThirdPartyInvoiceCallBackResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTaxInvoiceSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTaxInvoiceSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTaxInvoiceSyncAPIResponse)
	},
}

// GetAlibabaTaxInvoiceSyncAPIResponse 从 sync.Pool 获取 AlibabaTaxInvoiceSyncAPIResponse
func GetAlibabaTaxInvoiceSyncAPIResponse() *AlibabaTaxInvoiceSyncAPIResponse {
	return poolAlibabaTaxInvoiceSyncAPIResponse.Get().(*AlibabaTaxInvoiceSyncAPIResponse)
}

// ReleaseAlibabaTaxInvoiceSyncAPIResponse 将 AlibabaTaxInvoiceSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTaxInvoiceSyncAPIResponse(v *AlibabaTaxInvoiceSyncAPIResponse) {
	v.Reset()
	poolAlibabaTaxInvoiceSyncAPIResponse.Put(v)
}
