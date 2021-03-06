package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceBillEinvoiceListAPIResponse 扫码开票列表 API返回值
// alibaba.einvoice.bill.einvoice.list
//
// 扫码开票列表，包括用户扫二维码开票和结算单同步前的开票数据
type AlibabaEinvoiceBillEinvoiceListAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceBillEinvoiceListAPIResponseModel
}

// AlibabaEinvoiceBillEinvoiceListAPIResponseModel is 扫码开票列表 成功返回结果
type AlibabaEinvoiceBillEinvoiceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_bill_einvoice_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaEinvoiceBillEinvoiceListResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
