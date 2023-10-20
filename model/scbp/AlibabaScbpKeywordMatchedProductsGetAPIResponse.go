package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpKeywordMatchedProductsGetAPIResponse 查询和词匹配的推广产品 API返回值
// alibaba.scbp.keyword.matched.products.get
//
// 查询和词匹配的推广产品
type AlibabaScbpKeywordMatchedProductsGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpKeywordMatchedProductsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpKeywordMatchedProductsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpKeywordMatchedProductsGetAPIResponseModel).Reset()
}

// AlibabaScbpKeywordMatchedProductsGetAPIResponseModel is 查询和词匹配的推广产品 成功返回结果
type AlibabaScbpKeywordMatchedProductsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_keyword_matched_products_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 匹配的产品列表
	MachedProductList []TopMatchedProductDto `json:"mached_product_list,omitempty" xml:"mached_product_list>top_matched_product_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpKeywordMatchedProductsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MachedProductList = m.MachedProductList[:0]
}

var poolAlibabaScbpKeywordMatchedProductsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpKeywordMatchedProductsGetAPIResponse)
	},
}

// GetAlibabaScbpKeywordMatchedProductsGetAPIResponse 从 sync.Pool 获取 AlibabaScbpKeywordMatchedProductsGetAPIResponse
func GetAlibabaScbpKeywordMatchedProductsGetAPIResponse() *AlibabaScbpKeywordMatchedProductsGetAPIResponse {
	return poolAlibabaScbpKeywordMatchedProductsGetAPIResponse.Get().(*AlibabaScbpKeywordMatchedProductsGetAPIResponse)
}

// ReleaseAlibabaScbpKeywordMatchedProductsGetAPIResponse 将 AlibabaScbpKeywordMatchedProductsGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpKeywordMatchedProductsGetAPIResponse(v *AlibabaScbpKeywordMatchedProductsGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpKeywordMatchedProductsGetAPIResponse.Put(v)
}
