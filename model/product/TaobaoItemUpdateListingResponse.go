package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
一口价商品上架 APIResponse
taobao.item.update.listing

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateListingAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemUpdateListingResponse `json:"item_update_listing_response,omitempty"` 
    TaobaoItemUpdateListingResponse
}

/* model for simplify = false
type TaobaoItemUpdateListingResponse struct {

    // 上架后返回的商品信息：返回的结果就是:num_iid和modified
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoItemUpdateListingResponse struct {

    // 上架后返回的商品信息：返回的结果就是:num_iid和modified
    Item   *Item `json:"item,omitempty"`

}
