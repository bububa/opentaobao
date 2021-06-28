package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
taobao.item.update.listing天猫分流 APIResponse
taobao.item.update.listing.tmall

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateListingTmallAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemUpdateListingTmallResponse `json:"item_update_listing_tmall_response,omitempty"` 
    TaobaoItemUpdateListingTmallResponse
}

/* model for simplify = false
type TaobaoItemUpdateListingTmallResponse struct {

    // 上架后返回的商品信息：返回的结果就是:num_iid和modified
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoItemUpdateListingTmallResponse struct {

    // 上架后返回的商品信息：返回的结果就是:num_iid和modified
    Item   *Item `json:"item,omitempty"`

}
