package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionSchemaProductInstancePostAPIResponse aliexpress.solution.schema.product.instance.post API返回值
// aliexpress.solution.schema.product.instance.post
//
// Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
type AliexpressSolutionSchemaProductInstancePostAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionSchemaProductInstancePostAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionSchemaProductInstancePostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionSchemaProductInstancePostAPIResponseModel).Reset()
}

// AliexpressSolutionSchemaProductInstancePostAPIResponseModel is aliexpress.solution.schema.product.instance.post 成功返回结果
type AliexpressSolutionSchemaProductInstancePostAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_schema_product_instance_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result of the product posting
	Result *PostItemResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionSchemaProductInstancePostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionSchemaProductInstancePostAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionSchemaProductInstancePostAPIResponse)
	},
}

// GetAliexpressSolutionSchemaProductInstancePostAPIResponse 从 sync.Pool 获取 AliexpressSolutionSchemaProductInstancePostAPIResponse
func GetAliexpressSolutionSchemaProductInstancePostAPIResponse() *AliexpressSolutionSchemaProductInstancePostAPIResponse {
	return poolAliexpressSolutionSchemaProductInstancePostAPIResponse.Get().(*AliexpressSolutionSchemaProductInstancePostAPIResponse)
}

// ReleaseAliexpressSolutionSchemaProductInstancePostAPIResponse 将 AliexpressSolutionSchemaProductInstancePostAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionSchemaProductInstancePostAPIResponse(v *AliexpressSolutionSchemaProductInstancePostAPIResponse) {
	v.Reset()
	poolAliexpressSolutionSchemaProductInstancePostAPIResponse.Put(v)
}
