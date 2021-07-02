package aesolution

import (
	"encoding/xml"

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

// AliexpressSolutionSchemaProductInstancePostAPIResponseModel is aliexpress.solution.schema.product.instance.post 成功返回结果
type AliexpressSolutionSchemaProductInstancePostAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_schema_product_instance_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result of the product posting
	Result *PostItemResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}
