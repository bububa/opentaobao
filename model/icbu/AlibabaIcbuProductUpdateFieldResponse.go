package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品按字段更新 APIResponse
alibaba.icbu.product.update.field

按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
*/
type AlibabaIcbuProductUpdateFieldAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductUpdateFieldResponse `json:"alibaba_icbu_product_update_field_response,omitempty"` 
    AlibabaIcbuProductUpdateFieldResponse
}

/* model for simplify = false
type AlibabaIcbuProductUpdateFieldResponse struct {

    // 加密后的产品ID
    
    ProductId   string `json:"product_id,omitempty"`
    

}
*/

type AlibabaIcbuProductUpdateFieldResponse struct {

    // 加密后的产品ID
    ProductId   string `json:"product_id,omitempty"`

}
