package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词改价 APIRequest
alibaba.scbp.ad.keyword.price.update

关键词改价
*/
type AlibabaScbpAdKeywordPriceUpdateRequest struct {
    model.Params

    // 只能取ascci字符
    adKeyword   string 

    // 关键词价格单位元，一位小数
    priceStr   string 

}

func NewAlibabaScbpAdKeywordPriceUpdateRequest() *AlibabaScbpAdKeywordPriceUpdateRequest{
    return &AlibabaScbpAdKeywordPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.price.update"
}

func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordPriceUpdateRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetAdKeyword() string {
    return r.adKeyword
}

func (r *AlibabaScbpAdKeywordPriceUpdateRequest) SetPriceStr(priceStr string) error {
    r.priceStr = priceStr
    r.Set("price_str", priceStr)
    return nil
}

func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetPriceStr() string {
    return r.priceStr
}

