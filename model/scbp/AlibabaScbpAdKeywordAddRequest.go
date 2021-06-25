package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车加词 APIRequest
alibaba.scbp.ad.keyword.add

外贸直通车加词服务
*/
type AlibabaScbpAdKeywordAddRequest struct {
    model.Params

    // 加入的词
    adKeyword   string 

    // 词的出价
    priceStr   string 

    // 分组名
    tagName   string 

}

func NewAlibabaScbpAdKeywordAddRequest() *AlibabaScbpAdKeywordAddRequest{
    return &AlibabaScbpAdKeywordAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordAddRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.add"
}

func (r AlibabaScbpAdKeywordAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordAddRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

func (r AlibabaScbpAdKeywordAddRequest) GetAdKeyword() string {
    return r.adKeyword
}

func (r *AlibabaScbpAdKeywordAddRequest) SetPriceStr(priceStr string) error {
    r.priceStr = priceStr
    r.Set("price_str", priceStr)
    return nil
}

func (r AlibabaScbpAdKeywordAddRequest) GetPriceStr() string {
    return r.priceStr
}

func (r *AlibabaScbpAdKeywordAddRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

func (r AlibabaScbpAdKeywordAddRequest) GetTagName() string {
    return r.tagName
}

