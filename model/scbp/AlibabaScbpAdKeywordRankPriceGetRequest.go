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
    keyword   string
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
func (r *AlibabaScbpAdKeywordRankPriceGetRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r AlibabaScbpAdKeywordRankPriceGetRequest) GetKeyword() string {
    return r.keyword
}
