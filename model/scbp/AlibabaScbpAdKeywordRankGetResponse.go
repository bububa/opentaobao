package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取外贸直通车关键词预估排名 APIResponse
alibaba.scbp.ad.keyword.rank.get

获取外贸直通车关键词预估排名
*/
type AlibabaScbpAdKeywordRankGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdKeywordRankGetResponse `json:"alibaba_scbp_ad_keyword_rank_get_response,omitempty"` 
    AlibabaScbpAdKeywordRankGetResponse
}

/* model for simplify = false
type AlibabaScbpAdKeywordRankGetResponse struct {

    // 关键词的预估排名
    
    RankLocation   int64 `json:"rank_location,omitempty"`
    

}
*/

type AlibabaScbpAdKeywordRankGetResponse struct {

    // 关键词的预估排名
    RankLocation   int64 `json:"rank_location,omitempty"`

}
