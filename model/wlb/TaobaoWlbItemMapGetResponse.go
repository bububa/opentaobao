package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝商品ID查询商品映射关系 APIResponse
taobao.wlb.item.map.get

根据物流宝商品ID查询商品映射关系
*/
type TaobaoWlbItemMapGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_item_map_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 外部商品实体
    
    OutEntityItemList   []OutEntityItem `json:"out_entity_item_list,omitempty" xml:"