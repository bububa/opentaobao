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
    Response *TaobaoWlbItemCombinationGetResponse `json:"taobao_wlb_item_combination_get_response,omitempty"`
}

type TaobaoWlbItemCombinationGetResponse struct {

    // 组合子商品id列表
    ItemIdList   []Number `json:"item_id_list,omitempty"`

}
