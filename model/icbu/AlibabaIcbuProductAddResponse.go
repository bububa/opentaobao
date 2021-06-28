package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发布产品 APIResponse
alibaba.icbu.product.add

发布商品,支持sourcing/一口价商品，支持英文和多种语言原发商品
*/
type AlibabaIcbuProductAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductAddResponse `json:"alibaba_icbu_product_add_response,omitempty"` 
    AlibabaIcbuProductAddResponse
}

/* model for simplify = false
type AlibabaIcbuProductAddResponse struct {

    // 混淆后的产品ID
    
    ProductId   string `json:"product_id,omitempty"`
    

}
*/

type AlibabaIcbuProductAddResponse struct {

    // 混淆后的产品ID
    ProductId   string `json:"product_id,omitempty"`

}
