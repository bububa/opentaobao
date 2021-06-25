package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询和词匹配的推广产品 APIRequest
alibaba.scbp.keyword.matched.products.get

查询和词匹配的推广产品
*/
type AlibabaScbpKeywordMatchedProductsGetRequest struct {
    model.Params

    // 已购买的关键词
    adKeyword   string 

}

func NewAlibabaScbpKeywordMatchedProductsGetRequest() *AlibabaScbpKeywordMatchedProductsGetRequest{
    return &AlibabaScbpKeywordMatchedProductsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpKeywordMatchedProductsGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.keyword.matched.products.get"
}

func (r AlibabaScbpKeywordMatchedProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpKeywordMatchedProductsGetRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

func (r AlibabaScbpKeywordMatchedProductsGetRequest) GetAdKeyword() string {
    return r.adKeyword
}

