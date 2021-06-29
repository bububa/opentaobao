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
    adKeyword   string
    // 词的出价
    priceStr   string
    // 分组名
    tagName   string
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
func (r *AlibabaScbpAdKeywordAddRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpAdKeywordAddRequest) GetAdKeyword() string {
    return r.adKeyword
}
// PriceStr Setter
// 词的出价
func (r *AlibabaScbpAdKeywordAddRequest) SetPriceStr(priceStr string) error {
    r.priceStr = priceStr
    r.Set("price_str", priceStr)
    return nil
}

// PriceStr Getter
func (r AlibabaScbpAdKeywordAddRequest) GetPriceStr() string {
    return r.priceStr
}
// TagName Setter
// 分组名
func (r *AlibabaScbpAdKeywordAddRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

// TagName Getter
func (r AlibabaScbpAdKeywordAddRequest) GetTagName() string {
    return r.tagName
}
