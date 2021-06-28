package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id查询商品组合关系 APIResponse
taobao.wlb.item.combination.get

根据商品id查询商品组合关系
*/
type TaobaoWlbItemCombinationGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbItemCombinationGetResponse `json:"wlb_item_combination_get_response,omitempty"` 
    TaobaoWlbItemCombinationGetResponse
}

/* model for simplify = false
type TaobaoWlbItemCombinationGetResponse struct {

    // 组合子商品id列表
    
    ItemIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"item_id_list,omitempty"`
    

}
*/

type TaobaoWlbItemCombinationGetResponse struct {

    // 组合子商品id列表
    ItemIdList   []int64 `json:"item_id_list,omitempty"`

}
