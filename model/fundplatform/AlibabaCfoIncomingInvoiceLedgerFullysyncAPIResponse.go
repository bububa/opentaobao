package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfoincominginvoiceledgerfullysyncAPIResponse 票易通全量底账数据同步 API返回值
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
type AlibabacfoincominginvoiceledgerfullysyncAPIResponse struct {
	model.CommonResponse
	AlibabacfoincominginvoiceledgerfullysyncAPIResponseModel
}

// AlibabacfoincominginvoiceledgerfullysyncAPIResponseModel is 票易通全量底账数据同步 成功返回结果
type AlibabacfoincominginvoiceledgerfullysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_ledger_fullysync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步是否是否成功
	Result *AlibabacfoincominginvoiceledgerfullysyncResponse `json:"result,omitempty" xml:"result,omitempty"`
}
