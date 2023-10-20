package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductPostAPIResponse Product posting API API返回值
// aliexpress.solution.product.post
//
// Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
type AliexpressSolutionProductPostAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductPostAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionProductPostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionProductPostAPIResponseModel).Reset()
}

// AliexpressSolutionProductPostAPIResponseModel is Product posting API 成功返回结果
type AliexpressSolutionProductPostAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PostItemResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionProductPostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionProductPostAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionProductPostAPIResponse)
	},
}

// GetAliexpressSolutionProductPostAPIResponse 从 sync.Pool 获取 AliexpressSolutionProductPostAPIResponse
func GetAliexpressSolutionProductPostAPIResponse() *AliexpressSolutionProductPostAPIResponse {
	return poolAliexpressSolutionProductPostAPIResponse.Get().(*AliexpressSolutionProductPostAPIResponse)
}

// ReleaseAliexpressSolutionProductPostAPIResponse 将 AliexpressSolutionProductPostAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionProductPostAPIResponse(v *AliexpressSolutionProductPostAPIResponse) {
	v.Reset()
	poolAliexpressSolutionProductPostAPIResponse.Put(v)
}
