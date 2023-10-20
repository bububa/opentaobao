package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSearchAPIResponse 根据发票抬头搜索发票 API返回值
// alitrip.btrip.invoice.search
//
// 用户根据发票抬头搜索发票信息
type AlitripBtripInvoiceSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripInvoiceSearchAPIResponseModel).Reset()
}

// AlitripBtripInvoiceSearchAPIResponseModel is 根据发票抬头搜索发票 成功返回结果
type AlitripBtripInvoiceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripInvoiceSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripInvoiceSearchAPIResponse)
	},
}

// GetAlitripBtripInvoiceSearchAPIResponse 从 sync.Pool 获取 AlitripBtripInvoiceSearchAPIResponse
func GetAlitripBtripInvoiceSearchAPIResponse() *AlitripBtripInvoiceSearchAPIResponse {
	return poolAlitripBtripInvoiceSearchAPIResponse.Get().(*AlitripBtripInvoiceSearchAPIResponse)
}

// ReleaseAlitripBtripInvoiceSearchAPIResponse 将 AlitripBtripInvoiceSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripInvoiceSearchAPIResponse(v *AlitripBtripInvoiceSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripInvoiceSearchAPIResponse.Put(v)
}
