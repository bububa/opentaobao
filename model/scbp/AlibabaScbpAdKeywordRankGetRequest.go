package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取外贸直通车关键词预估排名 APIRequest
alibaba.scbp.ad.keyword.rank.get

获取外贸直通车关键词预估排名
*/
type AlibabaScbpAdKeywordRankGetRequest struct {
    model.Params

    // 查询预估排名的关键词
    keyword   string 

}

func NewAlibabaScbpAdKeywordRankGetRequest() *AlibabaScbpAdKeywordRankGetRequest{
    return &AlibabaScbpAdKeywordRankGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordRankGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.rank.get"
}

func (r AlibabaScbpAdKeywordRankGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordRankGetRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlibabaScbpAdKeywordRankGetRequest) GetKeyword() string {
    return r.keyword
}

