package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除单条商品 APIResponse
taobao.item.delete

删除单条商品
*/
type TaobaoItemDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemDeleteResponse `json:"item_delete_response,omitempty"` 
    TaobaoItemDeleteResponse
}

/* model for simplify = false
type TaobaoItemDeleteResponse struct {

    // 被删除商品的相关信息
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoItemDeleteResponse struct {

    // 被删除商品的相关信息
    Item   *Item `json:"item,omitempty"`

}
