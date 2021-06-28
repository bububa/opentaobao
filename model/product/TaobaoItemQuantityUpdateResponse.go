package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
宝贝/SKU库存修改 APIResponse
taobao.item.quantity.update

提供按照全量或增量形式修改宝贝/SKU库存的功能
*/
type TaobaoItemQuantityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemQuantityUpdateResponse `json:"item_quantity_update_response,omitempty"` 
    TaobaoItemQuantityUpdateResponse
}

/* model for simplify = false
type TaobaoItemQuantityUpdateResponse struct {

    // iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoItemQuantityUpdateResponse struct {

    // iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified
    Item   *Item `json:"item,omitempty"`

}
