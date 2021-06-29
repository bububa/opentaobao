package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益逆向扣回的查询接口 API请求
tmall.servicecenter.tp.funds.recover.query

服务商资金权益逆向扣回的查询接口
*/
type TmallServicecenterTpFundsRecoverQueryRequest struct {
    model.Params
    // query入参
    _query   *TpFundsRecoverQuery
}

// 初始化TmallServicecenterTpFundsRecoverQueryRequest对象
func NewTmallServicecenterTpFundsRecoverQueryRequest() *TmallServicecenterTpFundsRecoverQueryRequest{
    return &TmallServicecenterTpFundsRecoverQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTpFundsRecoverQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tp.funds.recover.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTpFundsRecoverQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// query入参
func (r *TmallServicecenterTpFundsRecoverQueryRequest) SetQuery(_query *TpFundsRecoverQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TmallServicecenterTpFundsRecoverQueryRequest) GetQuery() *TpFundsRecoverQuery {
    return r._query
}
