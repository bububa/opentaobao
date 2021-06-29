package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益发放的查询接口 API请求
tmall.servicecenter.tp.funds.send.query

服务商资金权益发放结果的查询接口
*/
type TmallServicecenterTpFundsSendQueryRequest struct {
    model.Params
    // 入参对象
    _query   *TpFundsSendQuery
}

// 初始化TmallServicecenterTpFundsSendQueryRequest对象
func NewTmallServicecenterTpFundsSendQueryRequest() *TmallServicecenterTpFundsSendQueryRequest{
    return &TmallServicecenterTpFundsSendQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTpFundsSendQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tp.funds.send.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTpFundsSendQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参对象
func (r *TmallServicecenterTpFundsSendQueryRequest) SetQuery(_query *TpFundsSendQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TmallServicecenterTpFundsSendQueryRequest) GetQuery() *TpFundsSendQuery {
    return r._query
}
