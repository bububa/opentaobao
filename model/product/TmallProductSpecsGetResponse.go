package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取产品的规格信息 APIResponse
tmall.product.specs.get

按product_id或品牌下载产品规格，返回一组的产品规格信息。
*/
type TmallProductSpecsGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallProductSpecsGetResponse `json:"tmall_product_specs_get_response,omitempty"` 
    TmallProductSpecsGetResponse
}

/* model for simplify = false
type TmallProductSpecsGetResponse struct {

    // 返回一组产品规格信息。
    
    ProductSpecs  struct {
        ProductSpec  []ProductSpec `json:"product_spec,omitempty"`
    } `json:"product_specs,omitempty"`
    

}
*/

type TmallProductSpecsGetResponse struct {

    // 返回一组产品规格信息。
    ProductSpecs   []ProductSpec `json:"product_specs,omitempty"`

}
