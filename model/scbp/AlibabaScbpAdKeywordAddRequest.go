package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车加词 API请求
alibaba.scbp.ad.keyword.add

外贸直通车加词服务
*/
type AlibabaScbpAdKeywordAddRequest struct {
    model.Params
    // 加入的词
    _adKeyword   string
    // 词的出价
    _priceStr   string
    // 分组名
    _tagName   string
}

// 初始化AlibabaScbpAdKeywordAddRequest对象
func NewAlibabaScbpAdKeywordAddRequest() *AlibabaScbpAdKeywordAddRequest{
    return &AlibabaScbpAdKeywordAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordAddRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdKeyword Setter
// 加入的词
func (r *AlibabaScbpAdKeywordAddRequest) SetAdKeyword(_adKeyword string) error {
    r._adKeyword = _adKeyword
    r.Set("ad_keyword", _adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpAdKeywordAddRequest) GetAdKeyword() string {
    return r._adKeyword
}
// PriceStr Setter
// 词的出价
func (r *AlibabaScbpAdKeywordAddRequest) SetPriceStr(_priceStr string) error {
    r._priceStr = _priceStr
    r.Set("price_str", _priceStr)
    return nil
}

// PriceStr Getter
func (r AlibabaScbpAdKeywordAddRequest) GetPriceStr() string {
    return r._priceStr
}
// TagName Setter
// 分组名
func (r *AlibabaScbpAdKeywordAddRequest) SetTagName(_tagName string) error {
    r._tagName = _tagName
    r.Set("tag_name", _tagName)
    return nil
}

// TagName Getter
func (r AlibabaScbpAdKeywordAddRequest) GetTagName() string {
    return r._tagName
}
