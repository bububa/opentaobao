package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id查询商品组合关系 APIResponse
taobao.wlb.item.combination.get

根据商品id查询商品组合关系
*/
type TaobaoWlbItemCombinationGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_item_combination_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 组合子商品id列表
    
    ItemIdList   []int64 `json:"item_id_list,omitempty" xml:"