package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词前五名排价 APIResponse
alibaba.scbp.ad.keyword.query.keyword.rank.price

查询关键词前五名排价
*/
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordQueryKeywordRankPriceResponse
}

type AlibabaScbpAdKeywordQueryKeywordRankPriceResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_query_keyword_rank_price_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回参数
    
    Result   *KeywordRankPriceDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
