package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝商品ID查询商品映射关系 APIResponse
taobao.wlb.item.map.get

根据物流宝商品ID查询商品映射关系
*/
type TaobaoWlbItemMapGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbItemMapGetResponse `json:"taobao_wlb_item_map_get_response,omitempty"`
}

type TaobaoWlbItemMapGetResponse struct {

    // 外部商品实体
    OutEntityItemList   []OutEntityItem `json:"out_entity_item_list,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
