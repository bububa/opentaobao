package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse 票易通全量底账数据同步 API返回值
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
type AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse struct {
	model.CommonResponse
	AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponseModel).Reset()
}

// AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponseModel is 票易通全量底账数据同步 成功返回结果
type AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_ledger_fullysync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步是否是否成功
	Result *AlibabaCfoIncomingInvoiceLedgerFullysyncResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse)
	},
}

// GetAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse 从 sync.Pool 获取 AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse
func GetAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse() *AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse {
	return poolAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse.Get().(*AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse)
}

// ReleaseAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse 将 AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse(v *AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse) {
	v.Reset()
	poolAlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse.Put(v)
}
