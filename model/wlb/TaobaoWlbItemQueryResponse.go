package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商品 APIResponse
taobao.wlb.item.query

根据状态、卖家、SKU等信息查询商品列表
*/
type TaobaoWlbItemQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbItemQueryResponse `json:"taobao_wlb_item_query_response,omitempty"`
}

type TaobaoWlbItemQueryResponse struct {

    // 结果总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 商品信息列表
    ItemList   []WlbItem `json:"item_list,omitempty"`

}
