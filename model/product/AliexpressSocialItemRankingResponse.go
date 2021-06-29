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
    AliexpressSocialItemRankingResponse
}

type AliexpressSocialItemRankingResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_item_ranking_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类型
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
