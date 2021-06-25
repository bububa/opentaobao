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
    Response *TaobaoWlbItemGetResponse `json:"taobao_wlb_item_get_response,omitempty"`
}

type TaobaoWlbItemGetResponse struct {

    // 商品信息
    Item   *WlbItem `json:"item,omitempty"`

}
