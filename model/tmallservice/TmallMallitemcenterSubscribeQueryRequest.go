package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务订购信息查询接口 API请求
tmall.mallitemcenter.subscribe.query

查询商家服务订购信息
*/
type TmallMallitemcenterSubscribeQueryAPIRequest struct {
    model.Params
    // 入参query
    _query   *Spb2bOderQuery
}

// 初始化TmallMallitemcenterSubscribeQueryAPIRequest对象
func NewTmallMallitemcenterSubscribeQueryRequest() *TmallMallitemcenterSubscribeQueryAPIRequest{
    return &TmallMallitemcenterSubscribeQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.subscribe.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参query
func (r *TmallMallitemcenterSubscribeQueryAPIRequest) SetQuery(_query *Spb2bOderQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetQuery() *Spb2bOderQuery {
    return r._query
}
