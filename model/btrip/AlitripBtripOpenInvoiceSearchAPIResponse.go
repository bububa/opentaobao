package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenInvoiceSearchAPIResponse 差旅申请用户搜索可用发票列表 API返回值
// alitrip.btrip.open.invoice.search
//
// 差旅申请用户搜索可用发票列表
type AlitripBtripOpenInvoiceSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenInvoiceSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenInvoiceSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenInvoiceSearchAPIResponseModel).Reset()
}

// AlitripBtripOpenInvoiceSearchAPIResponseModel is 差旅申请用户搜索可用发票列表 成功返回结果
type AlitripBtripOpenInvoiceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_invoice_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发票列表
	InvoiceList []OpenInvoiceDo `json:"invoice_list,omitempty" xml:"invoice_list>open_invoice_do,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenInvoiceSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InvoiceList = m.InvoiceList[:0]
	m.ResultMsg = ""
	m.ResultCode = 0
}

var poolAlitripBtripOpenInvoiceSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenInvoiceSearchAPIResponse)
	},
}

// GetAlitripBtripOpenInvoiceSearchAPIResponse 从 sync.Pool 获取 AlitripBtripOpenInvoiceSearchAPIResponse
func GetAlitripBtripOpenInvoiceSearchAPIResponse() *AlitripBtripOpenInvoiceSearchAPIResponse {
	return poolAlitripBtripOpenInvoiceSearchAPIResponse.Get().(*AlitripBtripOpenInvoiceSearchAPIResponse)
}

// ReleaseAlitripBtripOpenInvoiceSearchAPIResponse 将 AlitripBtripOpenInvoiceSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenInvoiceSearchAPIResponse(v *AlitripBtripOpenInvoiceSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenInvoiceSearchAPIResponse.Put(v)
}
