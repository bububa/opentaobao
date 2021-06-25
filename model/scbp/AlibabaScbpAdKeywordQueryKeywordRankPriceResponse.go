package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词前五名排价 APIResponse
alibaba.scbp.ad.keyword.query.keyword.rank.price

查询关键词前五名排价
*/
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordQueryKeywordRankPriceResponse `json:"alibaba_scbp_ad_keyword_query_keyword_rank_price_response,omitempty"`
}

type AlibabaScbpAdKeywordQueryKeywordRankPriceResponse struct {

    // 返回参数
    Result   *KeywordRankPriceDto `json:"result,omitempty"`

}
