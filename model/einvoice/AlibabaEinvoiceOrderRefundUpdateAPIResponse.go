package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceorderrefundupdateAPIResponse 回传订单退款审核结果 API返回值
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
type AlibabaeinvoiceorderrefundupdateAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoiceorderrefundupdateAPIResponseModel
}

// AlibabaeinvoiceorderrefundupdateAPIResponseModel is 回传订单退款审核结果 成功返回结果
type AlibabaeinvoiceorderrefundupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_order_refund_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
