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
type TmallMallitemcenterSubscribeQueryRequest struct {
    model.Params
    // 入参query
    query   *Spb2bOderQuery
}

// 初始化TmallMallitemcenterSubscribeQueryRequest对象
func NewTmallMallitemcenterSubscribeQueryRequest() *TmallMallitemcenterSubscribeQueryRequest{
    return &TmallMallitemcenterSubscribeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSubscribeQueryRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.subscribe.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSubscribeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参query
func (r *TmallMallitemcenterSubscribeQueryRequest) SetQuery(query *Spb2bOderQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r TmallMallitemcenterSubscribeQueryRequest) GetQuery() *Spb2bOderQuery {
    return r.query
}
