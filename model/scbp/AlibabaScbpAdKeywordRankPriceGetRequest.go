package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名排价 APIRequest
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价
*/
type AlibabaScbpAdKeywordRankPriceGetRequest struct {
    model.Params

    // 关键词
    keyword   string 

}

func NewAlibabaScbpAdKeywordRankPriceGetRequest() *AlibabaScbpAdKeywordRankPriceGetRequest{
    return &AlibabaScbpAdKeywordRankPriceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.rank.price.get"
}

func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordRankPriceGetRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetKeyword() string {
    return r.keyword
}

