package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统推荐 API请求
alibaba.scbp.reckeyword.sys.get

查询系统推荐词
*/
type AlibabaScbpReckeywordSysGetRequest struct {
    model.Params
    // RecKeywordQuery
    queryDto   *RecKeywordQuery
}

// 初始化AlibabaScbpReckeywordSysGetRequest对象
func NewAlibabaScbpReckeywordSysGetRequest() *AlibabaScbpReckeywordSysGetRequest{
    return &AlibabaScbpReckeywordSysGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpReckeywordSysGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.reckeyword.sys.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpReckeywordSysGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryDto Setter
// RecKeywordQuery
func (r *AlibabaScbpReckeywordSysGetRequest) SetQueryDto(queryDto *RecKeywordQuery) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

// QueryDto Getter
func (r AlibabaScbpReckeywordSysGetRequest) GetQueryDto() *RecKeywordQuery {
    return r.queryDto
}
