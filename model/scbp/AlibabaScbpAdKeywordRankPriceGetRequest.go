package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车关键词前五名排价 API请求
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价
*/
type AlibabaScbpAdKeywordRankPriceGetRequest struct {
    model.Params
    // 关键词
    _keyword   string
}

// 初始化AlibabaScbpAdKeywordRankPriceGetRequest对象
func NewAlibabaScbpAdKeywordRankPriceGetRequest() *AlibabaScbpAdKeywordRankPriceGetRequest{
    return &AlibabaScbpAdKeywordRankPriceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.rank.price.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 关键词
func (r *AlibabaScbpAdKeywordRankPriceGetRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetKeyword() string {
    return r._keyword
}
