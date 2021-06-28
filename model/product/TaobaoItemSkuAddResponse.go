package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加SKU APIResponse
taobao.item.sku.add

新增一个sku到num_iid指定的商品中 <br/>传入的iid所对应的商品必须属于当前会话的用户
*/
type TaobaoItemSkuAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemSkuAddResponse `json:"item_sku_add_response,omitempty"` 
    TaobaoItemSkuAddResponse
}

/* model for simplify = false
type TaobaoItemSkuAddResponse struct {

    // sku
    
    Sku  *struct {
        Sku  *Sku `json:"sku,omitempty"`
    } `json:"sku,omitempty"`
    

}
*/

type TaobaoItemSkuAddResponse struct {

    // sku
    Sku   *Sku `json:"sku,omitempty"`

}
