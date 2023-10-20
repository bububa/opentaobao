package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderInfoGetAPIResponse get order detail info API返回值
// aliexpress.solution.order.info.get
//
// get order detail info
type AliexpressSolutionOrderInfoGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionOrderInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionOrderInfoGetAPIResponseModel).Reset()
}

// AliexpressSolutionOrderInfoGetAPIResponseModel is get order detail info 成功返回结果
type AliexpressSolutionOrderInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionOrderInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionOrderInfoGetAPIResponse)
	},
}

// GetAliexpressSolutionOrderInfoGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionOrderInfoGetAPIResponse
func GetAliexpressSolutionOrderInfoGetAPIResponse() *AliexpressSolutionOrderInfoGetAPIResponse {
	return poolAliexpressSolutionOrderInfoGetAPIResponse.Get().(*AliexpressSolutionOrderInfoGetAPIResponse)
}

// ReleaseAliexpressSolutionOrderInfoGetAPIResponse 将 AliexpressSolutionOrderInfoGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionOrderInfoGetAPIResponse(v *AliexpressSolutionOrderInfoGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionOrderInfoGetAPIResponse.Put(v)
}
