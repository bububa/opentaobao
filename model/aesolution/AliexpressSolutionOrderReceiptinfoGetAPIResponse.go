package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderReceiptinfoGetAPIResponse Get Order Receipt Info API返回值
// aliexpress.solution.order.receiptinfo.get
//
// Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
type AliexpressSolutionOrderReceiptinfoGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionOrderReceiptinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderReceiptinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionOrderReceiptinfoGetAPIResponseModel).Reset()
}

// AliexpressSolutionOrderReceiptinfoGetAPIResponseModel is Get Order Receipt Info 成功返回结果
type AliexpressSolutionOrderReceiptinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_receiptinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderAddressDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderReceiptinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionOrderReceiptinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionOrderReceiptinfoGetAPIResponse)
	},
}

// GetAliexpressSolutionOrderReceiptinfoGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionOrderReceiptinfoGetAPIResponse
func GetAliexpressSolutionOrderReceiptinfoGetAPIResponse() *AliexpressSolutionOrderReceiptinfoGetAPIResponse {
	return poolAliexpressSolutionOrderReceiptinfoGetAPIResponse.Get().(*AliexpressSolutionOrderReceiptinfoGetAPIResponse)
}

// ReleaseAliexpressSolutionOrderReceiptinfoGetAPIResponse 将 AliexpressSolutionOrderReceiptinfoGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionOrderReceiptinfoGetAPIResponse(v *AliexpressSolutionOrderReceiptinfoGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionOrderReceiptinfoGetAPIResponse.Put(v)
}
