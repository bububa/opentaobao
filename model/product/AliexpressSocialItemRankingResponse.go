package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
社交排行榜 APIResponse
aliexpress.social.item.ranking

社交商品成交排行榜
*/
type AliexpressSocialItemRankingAPIResponse struct {
    model.CommonResponse
    // Response *AliexpressSocialItemRankingResponse `json:"aliexpress_social_item_ranking_response,omitempty"` 
    AliexpressSocialItemRankingResponse
}

/* model for simplify = false
type AliexpressSocialItemRankingResponse struct {

    // 返回包装类型
    
    Result  *struct {
        ItemPickPagingResult  *ItemPickPagingResult `json:"item_pick_paging_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AliexpressSocialItemRankingResponse struct {

    // 返回包装类型
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}
