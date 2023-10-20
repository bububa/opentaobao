package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionBatchProductDeleteAPIResponse aliexpress.solution.batch.product.delete API返回值
// aliexpress.solution.batch.product.delete
//
// Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
type AliexpressSolutionBatchProductDeleteAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionBatchProductDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionBatchProductDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionBatchProductDeleteAPIResponseModel).Reset()
}

// AliexpressSolutionBatchProductDeleteAPIResponseModel is aliexpress.solution.batch.product.delete 成功返回结果
type AliexpressSolutionBatchProductDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_batch_product_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionBatchProductDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAliexpressSolutionBatchProductDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionBatchProductDeleteAPIResponse)
	},
}

// GetAliexpressSolutionBatchProductDeleteAPIResponse 从 sync.Pool 获取 AliexpressSolutionBatchProductDeleteAPIResponse
func GetAliexpressSolutionBatchProductDeleteAPIResponse() *AliexpressSolutionBatchProductDeleteAPIResponse {
	return poolAliexpressSolutionBatchProductDeleteAPIResponse.Get().(*AliexpressSolutionBatchProductDeleteAPIResponse)
}

// ReleaseAliexpressSolutionBatchProductDeleteAPIResponse 将 AliexpressSolutionBatchProductDeleteAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionBatchProductDeleteAPIResponse(v *AliexpressSolutionBatchProductDeleteAPIResponse) {
	v.Reset()
	poolAliexpressSolutionBatchProductDeleteAPIResponse.Put(v)
}
