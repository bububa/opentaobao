package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益发放的查询接口 APIRequest
tmall.servicecenter.tp.funds.send.query

服务商资金权益发放结果的查询接口
*/
type TmallServicecenterTpFundsSendQueryRequest struct {
    model.Params

    // 入参对象
    query   *TpFundsSendQuery 

}

func NewTmallServicecenterTpFundsSendQueryRequest() *TmallServicecenterTpFundsSendQueryRequest{
    return &TmallServicecenterTpFundsSendQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterTpFundsSendQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tp.funds.send.query"
}

func (r TmallServicecenterTpFundsSendQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterTpFundsSendQueryRequest) SetQuery(query *TpFundsSendQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TmallServicecenterTpFundsSendQueryRequest) GetQuery() *TpFundsSendQuery {
    return r.query
}

