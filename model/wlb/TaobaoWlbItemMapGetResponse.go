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
    TaobaoWlbItemMapGetResponse
}

type TaobaoWlbItemMapGetResponse struct {
    XMLName xml.Name `xml:"wlb_item_map_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 外部商品实体
    
    OutEntityItemList   []OutEntityItem `json:"out_entity_item_list,omitempty" xml:"out_entity_item_list>out_entity_item,omitempty"`
    
    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
