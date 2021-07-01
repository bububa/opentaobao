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
type AlibabaScbpReckeywordSysGetAPIRequest struct {
    model.Params
    // RecKeywordQuery
    _queryDto   *RecKeywordQuery
}

// 初始化AlibabaScbpReckeywordSysGetAPIRequest对象
func NewAlibabaScbpReckeywordSysGetRequest() *AlibabaScbpReckeywordSysGetAPIRequest{
    return &AlibabaScbpReckeywordSysGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.reckeyword.sys.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryDto Setter
// RecKeywordQuery
func (r *AlibabaScbpReckeywordSysGetAPIRequest) SetQueryDto(_queryDto *RecKeywordQuery) error {
    r._queryDto = _queryDto
    r.Set("query_dto", _queryDto)
    return nil
}

// QueryDto Getter
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetQueryDto() *RecKeywordQuery {
    return r._queryDto
}
