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
type TmallServicecenterTpFundsSendQueryAPIRequest struct {
    model.Params
    // 入参对象
    _query   *TpFundsSendQuery
}

// 初始化TmallServicecenterTpFundsSendQueryAPIRequest对象
func NewTmallServicecenterTpFundsSendQueryRequest() *TmallServicecenterTpFundsSendQueryAPIRequest{
    return &TmallServicecenterTpFundsSendQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tp.funds.send.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参对象
func (r *TmallServicecenterTpFundsSendQueryAPIRequest) SetQuery(_query *TpFundsSendQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TmallServicecenterTpFundsSendQueryAPIRequest) GetQuery() *TpFundsSendQuery {
    return r._query
}
