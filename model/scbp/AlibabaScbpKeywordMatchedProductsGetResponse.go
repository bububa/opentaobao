package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询和词匹配的推广产品 APIResponse
alibaba.scbp.keyword.matched.products.get

查询和词匹配的推广产品
*/
type AlibabaScbpKeywordMatchedProductsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpKeywordMatchedProductsGetResponse `json:"alibaba_scbp_keyword_matched_products_get_response,omitempty"` 
    AlibabaScbpKeywordMatchedProductsGetResponse
}

/* model for simplify = false
type AlibabaScbpKeywordMatchedProductsGetResponse struct {

    // 匹配的产品列表
    
    MachedProductList  struct {
        TopMatchedProductDto  []TopMatchedProductDto `json:"top_matched_product_dto,omitempty"`
    } `json:"mached_product_list,omitempty"`
    

}
*/

type AlibabaScbpKeywordMatchedProductsGetResponse struct {

    // 匹配的产品列表
    MachedProductList   []TopMatchedProductDto `json:"mached_product_list,omitempty"`

}
