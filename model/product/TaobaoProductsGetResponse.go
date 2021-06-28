package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取产品列表 APIResponse
taobao.products.get

根据淘宝会员帐号搜索所有产品信息，推荐使用taobao.products.search
注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页
*/
type TaobaoProductsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoProductsGetResponse `json:"products_get_response,omitempty"` 
    TaobaoProductsGetResponse
}

/* model for simplify = false
type TaobaoProductsGetResponse struct {

    // 返回具体信息为入参fields请求的字段信息
    
    Products  struct {
        Product  []Product `json:"product,omitempty"`
    } `json:"products,omitempty"`
    

}
*/

type TaobaoProductsGetResponse struct {

    // 返回具体信息为入参fields请求的字段信息
    Products   []Product `json:"products,omitempty"`

}
