package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpKeywordMatchedProductsGetAPIResponse
查询和词匹配的推广产品 API返回值
alibaba.scbp.keyword.matched.products.get

查询和词匹配的推广产品 */
type AlibabaScbpKeywordMatchedProductsGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpKeywordMatchedProductsGetAPIResponseModel
}

// AlibabaScbpKeywordMatchedProductsGetAPIResponseModel is 查询和词匹配的推广产品 成功返回结果
type AlibabaScbpKeywordMatchedProductsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_keyword_matched_products_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 匹配的产品列表
	MachedProductList []TopMatchedProductDto `json:"mached_product_list,omitempty" xml:"mached_product_list>top_matched_product_dto,omitempty"`
}
