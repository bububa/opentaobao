package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获得单个商品详情 APIResponse
alibaba.icbu.product.get

获取商品详情
*/
type AlibabaIcbuProductGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductGetResponse `json:"alibaba_icbu_product_get_response,omitempty"` 
    AlibabaIcbuProductGetResponse
}

/* model for simplify = false
type AlibabaIcbuProductGetResponse struct {

    // 单个商品详情
    
    Product  *struct {
        AlibabaProductResponse  *AlibabaProductResponse `json:"alibaba_product_response,omitempty"`
    } `json:"product,omitempty"`
    

}
*/

type AlibabaIcbuProductGetResponse struct {

    // 单个商品详情
    Product   *AlibabaProductResponse `json:"product,omitempty"`

}
