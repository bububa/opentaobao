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
type AlibabaIdleIsvOrderQueryAPIRequest struct {
    model.Params
    // 系统自动生成
    _paramAppraiseIsvOrderQuery   *AppraiseIsvOrderQuery
}

// 初始化AlibabaIdleIsvOrderQueryAPIRequest对象
func NewAlibabaIdleIsvOrderQueryRequest() *AlibabaIdleIsvOrderQueryAPIRequest{
    return &AlibabaIdleIsvOrderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAppraiseIsvOrderQuery Setter
// 系统自动生成
func (r *AlibabaIdleIsvOrderQueryAPIRequest) SetParamAppraiseIsvOrderQuery(_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
    r._paramAppraiseIsvOrderQuery = _paramAppraiseIsvOrderQuery
    r.Set("param_appraise_isv_order_query", _paramAppraiseIsvOrderQuery)
    return nil
}

// ParamAppraiseIsvOrderQuery Getter
func (r AlibabaIdleIsvOrderQueryAPIRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
    return r._paramAppraiseIsvOrderQuery
}
