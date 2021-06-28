package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个关键词的推广组排名列表 APIResponse
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表
*/
type TaobaoSimbaToolsItemsTopGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaToolsItemsTopGetResponse `json:"simba_tools_items_top_get_response,omitempty"` 
    TaobaoSimbaToolsItemsTopGetResponse
}

/* model for simplify = false
type TaobaoSimbaToolsItemsTopGetResponse struct {

    // 推广组信息列表
    
    Rankeditems  struct {
        RankedItem  []RankedItem `json:"ranked_item,omitempty"`
    } `json:"rankeditems,omitempty"`
    

}
*/

type TaobaoSimbaToolsItemsTopGetResponse struct {

    // 推广组信息列表
    Rankeditems   []RankedItem `json:"rankeditems,omitempty"`

}
