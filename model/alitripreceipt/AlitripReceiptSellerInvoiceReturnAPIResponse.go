package alitripreceipt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripReceiptSellerInvoiceReturnAPIResponse 飞猪发票商家回调接口 API返回值
// alitrip.receipt.seller.invoice.return
//
// 飞猪发票回调接口
type AlitripReceiptSellerInvoiceReturnAPIResponse struct {
	model.CommonResponse
	AlitripReceiptSellerInvoiceReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripReceiptSellerInvoiceReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripReceiptSellerInvoiceReturnAPIResponseModel).Reset()
}

// AlitripReceiptSellerInvoiceReturnAPIResponseModel is 飞猪发票商家回调接口 成功返回结果
type AlitripReceiptSellerInvoiceReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_receipt_seller_invoice_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回数据
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripReceiptSellerInvoiceReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolAlitripReceiptSellerInvoiceReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripReceiptSellerInvoiceReturnAPIResponse)
	},
}

// GetAlitripReceiptSellerInvoiceReturnAPIResponse 从 sync.Pool 获取 AlitripReceiptSellerInvoiceReturnAPIResponse
func GetAlitripReceiptSellerInvoiceReturnAPIResponse() *AlitripReceiptSellerInvoiceReturnAPIResponse {
	return poolAlitripReceiptSellerInvoiceReturnAPIResponse.Get().(*AlitripReceiptSellerInvoiceReturnAPIResponse)
}

// ReleaseAlitripReceiptSellerInvoiceReturnAPIResponse 将 AlitripReceiptSellerInvoiceReturnAPIResponse 保存到 sync.Pool
func ReleaseAlitripReceiptSellerInvoiceReturnAPIResponse(v *AlitripReceiptSellerInvoiceReturnAPIResponse) {
	v.Reset()
	poolAlitripReceiptSellerInvoiceReturnAPIResponse.Put(v)
}
