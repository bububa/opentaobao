package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推荐词-词推词 API请求
alibaba.scbp.reckeyword.search

推荐词-词推词
*/
type AlibabaScbpReckeywordSearchRequest struct {
    model.Params
    // RecKeywordQuery
    queryDto   *RecKeywordQuery
}

// 初始化AlibabaScbpReckeywordSearchRequest对象
func NewAlibabaScbpReckeywordSearchRequest() *AlibabaScbpReckeywordSearchRequest{
    return &AlibabaScbpReckeywordSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpReckeywordSearchRequest) GetApiMethodName() string {
    return "alibaba.scbp.reckeyword.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpReckeywordSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryDto Setter
// RecKeywordQuery
func (r *AlibabaScbpReckeywordSearchRequest) SetQueryDto(queryDto *RecKeywordQuery) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

// QueryDto Getter
func (r AlibabaScbpReckeywordSearchRequest) GetQueryDto() *RecKeywordQuery {
    return r.queryDto
}
