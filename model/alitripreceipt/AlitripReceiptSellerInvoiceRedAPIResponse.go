package alitripreceipt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptSellerInvoiceRedAPIResponse 飞猪发票冲红 API返回值
// alitrip.receipt.seller.invoice.red
//
// 飞猪发票创建
type AlitripReceiptSellerInvoiceRedAPIResponse struct {
	model.CommonResponse
	AlitripReceiptSellerInvoiceRedAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripReceiptSellerInvoiceRedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripReceiptSellerInvoiceRedAPIResponseModel).Reset()
}

// AlitripReceiptSellerInvoiceRedAPIResponseModel is 飞猪发票冲红 成功返回结果
type AlitripReceiptSellerInvoiceRedAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_receipt_seller_invoice_red_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripReceiptSellerInvoiceRedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolAlitripReceiptSellerInvoiceRedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripReceiptSellerInvoiceRedAPIResponse)
	},
}

// GetAlitripReceiptSellerInvoiceRedAPIResponse 从 sync.Pool 获取 AlitripReceiptSellerInvoiceRedAPIResponse
func GetAlitripReceiptSellerInvoiceRedAPIResponse() *AlitripReceiptSellerInvoiceRedAPIResponse {
	return poolAlitripReceiptSellerInvoiceRedAPIResponse.Get().(*AlitripReceiptSellerInvoiceRedAPIResponse)
}

// ReleaseAlitripReceiptSellerInvoiceRedAPIResponse 将 AlitripReceiptSellerInvoiceRedAPIResponse 保存到 sync.Pool
func ReleaseAlitripReceiptSellerInvoiceRedAPIResponse(v *AlitripReceiptSellerInvoiceRedAPIResponse) {
	v.Reset()
	poolAlitripReceiptSellerInvoiceRedAPIResponse.Put(v)
}
