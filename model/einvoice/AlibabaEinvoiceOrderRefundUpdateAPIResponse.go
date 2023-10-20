package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceOrderRefundUpdateAPIResponse 回传订单退款审核结果 API返回值
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
type AlibabaEinvoiceOrderRefundUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceOrderRefundUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceOrderRefundUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceOrderRefundUpdateAPIResponseModel).Reset()
}

// AlibabaEinvoiceOrderRefundUpdateAPIResponseModel is 回传订单退款审核结果 成功返回结果
type AlibabaEinvoiceOrderRefundUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_order_refund_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceOrderRefundUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaEinvoiceOrderRefundUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceOrderRefundUpdateAPIResponse)
	},
}

// GetAlibabaEinvoiceOrderRefundUpdateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceOrderRefundUpdateAPIResponse
func GetAlibabaEinvoiceOrderRefundUpdateAPIResponse() *AlibabaEinvoiceOrderRefundUpdateAPIResponse {
	return poolAlibabaEinvoiceOrderRefundUpdateAPIResponse.Get().(*AlibabaEinvoiceOrderRefundUpdateAPIResponse)
}

// ReleaseAlibabaEinvoiceOrderRefundUpdateAPIResponse 将 AlibabaEinvoiceOrderRefundUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceOrderRefundUpdateAPIResponse(v *AlibabaEinvoiceOrderRefundUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceOrderRefundUpdateAPIResponse.Put(v)
}
