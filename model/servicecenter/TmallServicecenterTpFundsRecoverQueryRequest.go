package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益逆向扣回的查询接口 APIRequest
tmall.servicecenter.tp.funds.recover.query

服务商资金权益逆向扣回的查询接口
*/
type TmallServicecenterTpFundsRecoverQueryRequest struct {
    model.Params

    // query入参
    query   *TpFundsRecoverQuery 

}

func NewTmallServicecenterTpFundsRecoverQueryRequest() *TmallServicecenterTpFundsRecoverQueryRequest{
    return &TmallServicecenterTpFundsRecoverQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterTpFundsRecoverQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tp.funds.recover.query"
}

func (r TmallServicecenterTpFundsRecoverQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterTpFundsRecoverQueryRequest) SetQuery(query *TpFundsRecoverQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TmallServicecenterTpFundsRecoverQueryRequest) GetQuery() *TpFundsRecoverQuery {
    return r.query
}

