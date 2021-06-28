package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询和词匹配的推广产品 APIResponse
alibaba.scbp.keyword.matched.products.get

查询和词匹配的推广产品
*/
type AlibabaScbpKeywordMatchedProductsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_keyword_matched_products_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 匹配的产品列表
    
    MachedProductList   []TopMatchedProductDto `json:"mached_product_list,omitempty" xml:"