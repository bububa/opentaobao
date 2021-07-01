package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id查询商品组合关系 API返回值 
taobao.wlb.item.combination.get

根据商品id查询商品组合关系
*/
type TaobaoWlbItemCombinationGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbItemCombinationGetAPIResponseModel
}

// 根据商品id查询商品组合关系 成功返回结果
type TaobaoWlbItemCombinationGetAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_item_combination_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 组合子商品id列表
    ItemIdList   []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
}
