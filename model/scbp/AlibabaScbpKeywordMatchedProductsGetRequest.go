package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询和词匹配的推广产品 API请求
alibaba.scbp.keyword.matched.products.get

查询和词匹配的推广产品
*/
type AlibabaScbpKeywordMatchedProductsGetRequest struct {
    model.Params
    // 已购买的关键词
    _adKeyword   string
}

// 初始化AlibabaScbpKeywordMatchedProductsGetRequest对象
func NewAlibabaScbpKeywordMatchedProductsGetRequest() *AlibabaScbpKeywordMatchedProductsGetRequest{
    return &AlibabaScbpKeywordMatchedProductsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpKeywordMatchedProductsGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.keyword.matched.products.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpKeywordMatchedProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdKeyword Setter
// 已购买的关键词
func (r *AlibabaScbpKeywordMatchedProductsGetRequest) SetAdKeyword(_adKeyword string) error {
    r._adKeyword = _adKeyword
    r.Set("ad_keyword", _adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpKeywordMatchedProductsGetRequest) GetAdKeyword() string {
    return r._adKeyword
}
