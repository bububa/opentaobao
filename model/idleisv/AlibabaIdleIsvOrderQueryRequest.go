package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货订单查询 API请求
alibaba.idle.isv.order.query

服务商ISV，根据订单号，查询闲鱼订单信息
*/
type AlibabaIdleIsvOrderQueryRequest struct {
    model.Params
    // 系统自动生成
    _paramAppraiseIsvOrderQuery   *AppraiseIsvOrderQuery
}

// 初始化AlibabaIdleIsvOrderQueryRequest对象
func NewAlibabaIdleIsvOrderQueryRequest() *AlibabaIdleIsvOrderQueryRequest{
    return &AlibabaIdleIsvOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAppraiseIsvOrderQuery Setter
// 系统自动生成
func (r *AlibabaIdleIsvOrderQueryRequest) SetParamAppraiseIsvOrderQuery(_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
    r._paramAppraiseIsvOrderQuery = _paramAppraiseIsvOrderQuery
    r.Set("param_appraise_isv_order_query", _paramAppraiseIsvOrderQuery)
    return nil
}

// ParamAppraiseIsvOrderQuery Getter
func (r AlibabaIdleIsvOrderQueryRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
    return r._paramAppraiseIsvOrderQuery
}
