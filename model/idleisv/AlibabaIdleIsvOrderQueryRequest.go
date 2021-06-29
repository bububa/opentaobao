package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货订单查询 APIRequest
alibaba.idle.isv.order.query

服务商ISV，根据订单号，查询闲鱼订单信息
*/
type AlibabaIdleIsvOrderQueryRequest struct {
    model.Params

    // 系统自动生成
    paramAppraiseIsvOrderQuery   *AppraiseIsvOrderQuery 

}

func NewAlibabaIdleIsvOrderQueryRequest() *AlibabaIdleIsvOrderQueryRequest{
    return &AlibabaIdleIsvOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.query"
}

func (r AlibabaIdleIsvOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvOrderQueryRequest) SetParamAppraiseIsvOrderQuery(paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
    r.paramAppraiseIsvOrderQuery = paramAppraiseIsvOrderQuery
    r.Set("param_appraise_isv_order_query", paramAppraiseIsvOrderQuery)
    return nil
}

func (r AlibabaIdleIsvOrderQueryRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
    return r.paramAppraiseIsvOrderQuery
}

