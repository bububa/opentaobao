package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductSchemaGetAPIResponse get product schema API返回值
// aliexpress.solution.product.schema.get
//
// provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
type AliexpressSolutionProductSchemaGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionProductSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionProductSchemaGetAPIResponseModel).Reset()
}

// AliexpressSolutionProductSchemaGetAPIResponseModel is get product schema 成功返回结果
type AliexpressSolutionProductSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ProductSchemaDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionProductSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionProductSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionProductSchemaGetAPIResponse)
	},
}

// GetAliexpressSolutionProductSchemaGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionProductSchemaGetAPIResponse
func GetAliexpressSolutionProductSchemaGetAPIResponse() *AliexpressSolutionProductSchemaGetAPIResponse {
	return poolAliexpressSolutionProductSchemaGetAPIResponse.Get().(*AliexpressSolutionProductSchemaGetAPIResponse)
}

// ReleaseAliexpressSolutionProductSchemaGetAPIResponse 将 AliexpressSolutionProductSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionProductSchemaGetAPIResponse(v *AliexpressSolutionProductSchemaGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionProductSchemaGetAPIResponse.Put(v)
}
