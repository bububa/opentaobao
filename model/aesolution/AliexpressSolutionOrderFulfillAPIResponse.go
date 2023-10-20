package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderFulfillAPIResponse fulfill order API返回值
// aliexpress.solution.order.fulfill
//
// fulfill order for seller
type AliexpressSolutionOrderFulfillAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionOrderFulfillAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderFulfillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionOrderFulfillAPIResponseModel).Reset()
}

// AliexpressSolutionOrderFulfillAPIResponseModel is fulfill order 成功返回结果
type AliexpressSolutionOrderFulfillAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_fulfill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderFulfillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionOrderFulfillAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionOrderFulfillAPIResponse)
	},
}

// GetAliexpressSolutionOrderFulfillAPIResponse 从 sync.Pool 获取 AliexpressSolutionOrderFulfillAPIResponse
func GetAliexpressSolutionOrderFulfillAPIResponse() *AliexpressSolutionOrderFulfillAPIResponse {
	return poolAliexpressSolutionOrderFulfillAPIResponse.Get().(*AliexpressSolutionOrderFulfillAPIResponse)
}

// ReleaseAliexpressSolutionOrderFulfillAPIResponse 将 AliexpressSolutionOrderFulfillAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionOrderFulfillAPIResponse(v *AliexpressSolutionOrderFulfillAPIResponse) {
	v.Reset()
	poolAliexpressSolutionOrderFulfillAPIResponse.Put(v)
}
