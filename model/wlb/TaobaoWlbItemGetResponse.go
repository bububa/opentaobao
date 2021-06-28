package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID获取商品信息 APIResponse
taobao.wlb.item.get

根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
*/
type TaobaoWlbItemGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbItemGetResponse `json:"wlb_item_get_response,omitempty"` 
    TaobaoWlbItemGetResponse
}

/* model for simplify = false
type TaobaoWlbItemGetResponse struct {

    // 商品信息
    
    Item  *struct {
        WlbItem  *WlbItem `json:"wlb_item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

type TaobaoWlbItemGetResponse struct {

    // 商品信息
    Item   *WlbItem `json:"item,omitempty"`

}
