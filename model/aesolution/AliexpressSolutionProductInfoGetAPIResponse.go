package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductInfoGetAPIResponse Get Single Product Info API返回值
// aliexpress.solution.product.info.get
//
// Get Single Product Info
type AliexpressSolutionProductInfoGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionProductInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionProductInfoGetAPIResponseModel).Reset()
}

// AliexpressSolutionProductInfoGetAPIResponseModel is Get Single Product Info 成功返回结果
type AliexpressSolutionProductInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *GlobalAeopFindProductResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionProductInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionProductInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionProductInfoGetAPIResponse)
	},
}

// GetAliexpressSolutionProductInfoGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionProductInfoGetAPIResponse
func GetAliexpressSolutionProductInfoGetAPIResponse() *AliexpressSolutionProductInfoGetAPIResponse {
	return poolAliexpressSolutionProductInfoGetAPIResponse.Get().(*AliexpressSolutionProductInfoGetAPIResponse)
}

// ReleaseAliexpressSolutionProductInfoGetAPIResponse 将 AliexpressSolutionProductInfoGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionProductInfoGetAPIResponse(v *AliexpressSolutionProductInfoGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionProductInfoGetAPIResponse.Put(v)
}
