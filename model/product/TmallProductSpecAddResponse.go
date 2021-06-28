package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加产品规格 APIResponse
tmall.product.spec.add

增加产品规格
*/
type TmallProductSpecAddAPIResponse struct {
    model.CommonResponse
    // Response *TmallProductSpecAddResponse `json:"tmall_product_spec_add_response,omitempty"` 
    TmallProductSpecAddResponse
}

/* model for simplify = false
type TmallProductSpecAddResponse struct {

    // 产品规格对象
    
    ProductSpec  *struct {
        ProductSpec  *ProductSpec `json:"product_spec,omitempty"`
    } `json:"product_spec,omitempty"`
    

}
*/

type TmallProductSpecAddResponse struct {

    // 产品规格对象
    ProductSpec   *ProductSpec `json:"product_spec,omitempty"`

}
