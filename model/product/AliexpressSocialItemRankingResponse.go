package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
社交排行榜 APIResponse
aliexpress.social.item.ranking

社交商品成交排行榜
*/
type AliexpressSocialItemRankingAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"aliexpress_social_item_ranking_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回包装类型
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"