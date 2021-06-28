package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名排价 APIResponse
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价
*/
type AlibabaScbpAdKeywordRankPriceGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordRankPriceGetResponse
}

type AlibabaScbpAdKeywordRankPriceGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_rank_price_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词前五名排价
    
    RankPriceList   []string `json:"rank_price_list,omitempty" xml:"rank_price_list>string,omitempty"`
    
    
}
