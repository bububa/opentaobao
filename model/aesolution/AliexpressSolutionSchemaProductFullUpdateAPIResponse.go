package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionSchemaProductFullUpdateAPIResponse aliexpress.solution.schema.product.full.update API返回值
// aliexpress.solution.schema.product.full.update
//
// Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
type AliexpressSolutionSchemaProductFullUpdateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionSchemaProductFullUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionSchemaProductFullUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionSchemaProductFullUpdateAPIResponseModel).Reset()
}

// AliexpressSolutionSchemaProductFullUpdateAPIResponseModel is aliexpress.solution.schema.product.full.update 成功返回结果
type AliexpressSolutionSchemaProductFullUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_schema_product_full_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Product id that has been updated.
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionSchemaProductFullUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductId = 0
}

var poolAliexpressSolutionSchemaProductFullUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionSchemaProductFullUpdateAPIResponse)
	},
}

// GetAliexpressSolutionSchemaProductFullUpdateAPIResponse 从 sync.Pool 获取 AliexpressSolutionSchemaProductFullUpdateAPIResponse
func GetAliexpressSolutionSchemaProductFullUpdateAPIResponse() *AliexpressSolutionSchemaProductFullUpdateAPIResponse {
	return poolAliexpressSolutionSchemaProductFullUpdateAPIResponse.Get().(*AliexpressSolutionSchemaProductFullUpdateAPIResponse)
}

// ReleaseAliexpressSolutionSchemaProductFullUpdateAPIResponse 将 AliexpressSolutionSchemaProductFullUpdateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionSchemaProductFullUpdateAPIResponse(v *AliexpressSolutionSchemaProductFullUpdateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionSchemaProductFullUpdateAPIResponse.Put(v)
}
