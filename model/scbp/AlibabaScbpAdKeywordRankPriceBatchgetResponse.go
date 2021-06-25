package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名批量排价 APIResponse
alibaba.scbp.ad.keyword.rank.price.batchget

外贸直通车关键词前五名批量排价
*/
type AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordRankPriceBatchgetResponse `json:"alibaba_scbp_ad_keyword_rank_price_batchget_response,omitempty"`
}

type AlibabaScbpAdKeywordRankPriceBatchgetResponse struct {

    // 返回结果
    ResultList   []KeywordRankPriceDTO `json:"result_list,omitempty"`

}
