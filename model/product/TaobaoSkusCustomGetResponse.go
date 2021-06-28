package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据外部ID取商品SKU APIResponse
taobao.skus.custom.get

跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku <br/>这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
*/
type TaobaoSkusCustomGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSkusCustomGetResponse `json:"skus_custom_get_response,omitempty"` 
    TaobaoSkusCustomGetResponse
}

/* model for simplify = false
type TaobaoSkusCustomGetResponse struct {

    // Sku对象，具体字段以fields决定
    
    Skus  struct {
        Sku  []Sku `json:"sku,omitempty"`
    } `json:"skus,omitempty"`
    

}
*/

type TaobaoSkusCustomGetResponse struct {

    // Sku对象，具体字段以fields决定
    Skus   []Sku `json:"skus,omitempty"`

}
