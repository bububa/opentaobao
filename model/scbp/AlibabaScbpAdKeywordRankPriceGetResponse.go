package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名排价 APIResponse
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价
*/
type AlibabaScbpAdKeywordRankPriceGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordRankPriceGetResponse `json:"alibaba_scbp_ad_keyword_rank_price_get_response,omitempty"` 
    AlibabaScbpAdKeywordRankPriceGetResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordRankPriceGetResponse struct {

    // 关键词前五名排价
    
    RankPriceList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"rank_price_list,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordRankPriceGetResponse struct {

    // 关键词前五名排价
    RankPriceList   []string `json:"rank_price_list,omitempty"`

}
