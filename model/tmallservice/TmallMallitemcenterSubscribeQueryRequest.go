package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务订购信息查询接口 APIRequest
tmall.mallitemcenter.subscribe.query

查询商家服务订购信息
*/
type TmallMallitemcenterSubscribeQueryRequest struct {
    model.Params

    // 入参query
    query   *Spb2bOderQuery 

}

func NewTmallMallitemcenterSubscribeQueryRequest() *TmallMallitemcenterSubscribeQueryRequest{
    return &TmallMallitemcenterSubscribeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMallitemcenterSubscribeQueryRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.subscribe.query"
}

func (r TmallMallitemcenterSubscribeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMallitemcenterSubscribeQueryRequest) SetQuery(query *Spb2bOderQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TmallMallitemcenterSubscribeQueryRequest) GetQuery() *Spb2bOderQuery {
    return r.query
}

