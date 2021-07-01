package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车删除关键词 API请求
alibaba.scbp.ad.keyword.delete

外贸直通车删除关键词
*/
type AlibabaScbpAdKeywordDeleteAPIRequest struct {
    model.Params
    // 要删除的关键词
    _adKeyword   string
}

// 初始化AlibabaScbpAdKeywordDeleteAPIRequest对象
func NewAlibabaScbpAdKeywordDeleteRequest() *AlibabaScbpAdKeywordDeleteAPIRequest{
    return &AlibabaScbpAdKeywordDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdKeyword Setter
// 要删除的关键词
func (r *AlibabaScbpAdKeywordDeleteAPIRequest) SetAdKeyword(_adKeyword string) error {
    r._adKeyword = _adKeyword
    r.Set("ad_keyword", _adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetAdKeyword() string {
    return r._adKeyword
}
