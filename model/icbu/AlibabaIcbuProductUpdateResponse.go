package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改商品 APIResponse
alibaba.icbu.product.update

修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
*/
type AlibabaIcbuProductUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductUpdateResponse `json:"alibaba_icbu_product_update_response,omitempty"` 
    AlibabaIcbuProductUpdateResponse
}

/* model for simplify = false
type AlibabaIcbuProductUpdateResponse struct {

    // 加密后的产品ID
    
    ProductId   string `json:"product_id,omitempty"`
    

}
*/

type AlibabaIcbuProductUpdateResponse struct {

    // 加密后的产品ID
    ProductId   string `json:"product_id,omitempty"`

}
