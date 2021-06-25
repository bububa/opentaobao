package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车删除关键词 APIRequest
alibaba.scbp.ad.keyword.delete

外贸直通车删除关键词
*/
type AlibabaScbpAdKeywordDeleteRequest struct {
    model.Params

    // 要删除的关键词
    adKeyword   string 

}

func NewAlibabaScbpAdKeywordDeleteRequest() *AlibabaScbpAdKeywordDeleteRequest{
    return &AlibabaScbpAdKeywordDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordDeleteRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.delete"
}

func (r AlibabaScbpAdKeywordDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordDeleteRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

func (r AlibabaScbpAdKeywordDeleteRequest) GetAdKeyword() string {
    return r.adKeyword
}

