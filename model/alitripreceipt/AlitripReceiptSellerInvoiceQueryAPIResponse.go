package alitripreceipt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripReceiptSellerInvoiceQueryAPIResponse
飞猪发票查询 API返回值
alitrip.receipt.seller.invoice.query

飞猪发票查询 */
type AlitripReceiptSellerInvoiceQueryAPIResponse struct {
	model.CommonResponse
	AlitripReceiptSellerInvoiceQueryAPIResponseModel
}

// AlitripReceiptSellerInvoiceQueryAPIResponseModel is 飞猪发票查询 成功返回结果
type AlitripReceiptSellerInvoiceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_receipt_seller_invoice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回发票数据
	ReceiptDOs string `json:"receipt_d_os,omitempty" xml:"receipt_d_os,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
