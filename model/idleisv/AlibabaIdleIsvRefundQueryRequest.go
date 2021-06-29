package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货交易订单退款信息查询 APIRequest
alibaba.idle.isv.refund.query

闲鱼服务商交易订单退款信息查询-单个退款查询
*/
type AlibabaIdleIsvRefundQueryRequest struct {
    model.Params

    // 系统自动生成
    paramAppraiseIsvOrderQuery   *AppraiseIsvOrderQuery 

}

func NewAlibabaIdleIsvRefundQueryRequest() *AlibabaIdleIsvRefundQueryRequest{
    return &AlibabaIdleIsvRefundQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvRefundQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.refund.query"
}

func (r AlibabaIdleIsvRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvRefundQueryRequest) SetParamAppraiseIsvOrderQuery(paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
    r.paramAppraiseIsvOrderQuery = paramAppraiseIsvOrderQuery
    r.Set("param_appraise_isv_order_query", paramAppraiseIsvOrderQuery)
    return nil
}

func (r AlibabaIdleIsvRefundQueryRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
    return r.paramAppraiseIsvOrderQuery
}

