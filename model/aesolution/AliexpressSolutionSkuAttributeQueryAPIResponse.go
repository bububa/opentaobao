package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionSkuAttributeQueryAPIResponse Query the sku attribute information belonged to a specific category API返回值
// aliexpress.solution.sku.attribute.query
//
// Query the sku attribute information belonged to a specific category, customized for oversea merchants.
type AliexpressSolutionSkuAttributeQueryAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionSkuAttributeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionSkuAttributeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionSkuAttributeQueryAPIResponseModel).Reset()
}

// AliexpressSolutionSkuAttributeQueryAPIResponseModel is Query the sku attribute information belonged to a specific category 成功返回结果
type AliexpressSolutionSkuAttributeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_sku_attribute_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SkuAttributeInfoQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionSkuAttributeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionSkuAttributeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionSkuAttributeQueryAPIResponse)
	},
}

// GetAliexpressSolutionSkuAttributeQueryAPIResponse 从 sync.Pool 获取 AliexpressSolutionSkuAttributeQueryAPIResponse
func GetAliexpressSolutionSkuAttributeQueryAPIResponse() *AliexpressSolutionSkuAttributeQueryAPIResponse {
	return poolAliexpressSolutionSkuAttributeQueryAPIResponse.Get().(*AliexpressSolutionSkuAttributeQueryAPIResponse)
}

// ReleaseAliexpressSolutionSkuAttributeQueryAPIResponse 将 AliexpressSolutionSkuAttributeQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionSkuAttributeQueryAPIResponse(v *AliexpressSolutionSkuAttributeQueryAPIResponse) {
	v.Reset()
	poolAliexpressSolutionSkuAttributeQueryAPIResponse.Put(v)
}
