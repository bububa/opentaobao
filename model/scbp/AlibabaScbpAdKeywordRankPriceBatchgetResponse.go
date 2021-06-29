package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名批量排价 API返回值 
alibaba.scbp.ad.keyword.rank.price.batchget

外贸直通车关键词前五名批量排价
*/
type AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordRankPriceBatchgetResponse
}

// 外贸直通车关键词前五名批量排价 成功返回结果
type AlibabaScbpAdKeywordRankPriceBatchgetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_rank_price_batchget_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    ResultList   []KeywordRankPriceDTO `json:"result_list,omitempty" xml:"result_list>keyword_rank_price_dto,omitempty"`
}
