package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceGetAPIResponse 获取用户可用发票列表 API返回值
// alitrip.btrip.invoice.get
//
// 差旅申请用户获取可用发票列表
type AlitripBtripInvoiceGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripInvoiceGetAPIResponseModel).Reset()
}

// AlitripBtripInvoiceGetAPIResponseModel is 获取用户可用发票列表 成功返回结果
type AlitripBtripInvoiceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripInvoiceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripInvoiceGetAPIResponse)
	},
}

// GetAlitripBtripInvoiceGetAPIResponse 从 sync.Pool 获取 AlitripBtripInvoiceGetAPIResponse
func GetAlitripBtripInvoiceGetAPIResponse() *AlitripBtripInvoiceGetAPIResponse {
	return poolAlitripBtripInvoiceGetAPIResponse.Get().(*AlitripBtripInvoiceGetAPIResponse)
}

// ReleaseAlitripBtripInvoiceGetAPIResponse 将 AlitripBtripInvoiceGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripInvoiceGetAPIResponse(v *AlitripBtripInvoiceGetAPIResponse) {
	v.Reset()
	poolAlitripBtripInvoiceGetAPIResponse.Put(v)
}
