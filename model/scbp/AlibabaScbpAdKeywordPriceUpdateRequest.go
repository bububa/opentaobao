package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词改价 API请求
alibaba.scbp.ad.keyword.price.update

关键词改价
*/
type AlibabaScbpAdKeywordPriceUpdateRequest struct {
    model.Params
    // 只能取ascci字符
    _adKeyword   string
    // 关键词价格单位元，一位小数
    _priceStr   string
}

// 初始化AlibabaScbpAdKeywordPriceUpdateRequest对象
func NewAlibabaScbpAdKeywordPriceUpdateRequest() *AlibabaScbpAdKeywordPriceUpdateRequest{
    return &AlibabaScbpAdKeywordPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.price.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdKeyword Setter
// 只能取ascci字符
func (r *AlibabaScbpAdKeywordPriceUpdateRequest) SetAdKeyword(_adKeyword string) error {
    r._adKeyword = _adKeyword
    r.Set("ad_keyword", _adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetAdKeyword() string {
    return r._adKeyword
}
// PriceStr Setter
// 关键词价格单位元，一位小数
func (r *AlibabaScbpAdKeywordPriceUpdateRequest) SetPriceStr(_priceStr string) error {
    r._priceStr = _priceStr
    r.Set("price_str", _priceStr)
    return nil
}

// PriceStr Getter
func (r AlibabaScbpAdKeywordPriceUpdateRequest) GetPriceStr() string {
    return r._priceStr
}
