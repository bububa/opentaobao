package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车查询关键词 API请求
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词
*/
type AlibabaScbpAdKeywordGetRequest struct {
    model.Params
    // KeywordQuery
    queryDto   *KeywordQuery
}

// 初始化AlibabaScbpAdKeywordGetRequest对象
func NewAlibabaScbpAdKeywordGetRequest() *AlibabaScbpAdKeywordGetRequest{
    return &AlibabaScbpAdKeywordGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryDto Setter
// KeywordQuery
func (r *AlibabaScbpAdKeywordGetRequest) SetQueryDto(queryDto *KeywordQuery) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

// QueryDto Getter
func (r AlibabaScbpAdKeywordGetRequest) GetQueryDto() *KeywordQuery {
    return r.queryDto
}
