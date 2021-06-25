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
    Response *AliexpressSocialItemRankingResponse `json:"aliexpress_social_item_ranking_response,omitempty"`
}

type AliexpressSocialItemRankingResponse struct {

    // 返回包装类型
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}
