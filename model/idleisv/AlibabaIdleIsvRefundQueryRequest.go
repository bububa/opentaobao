package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货交易订单退款信息查询 API请求
alibaba.idle.isv.refund.query

闲鱼服务商交易订单退款信息查询-单个退款查询
*/
type AlibabaIdleIsvRefundQueryRequest struct {
    model.Params
    // 系统自动生成
    _paramAppraiseIsvOrderQuery   *AppraiseIsvOrderQuery
}

// 初始化AlibabaIdleIsvRefundQueryRequest对象
func NewAlibabaIdleIsvRefundQueryRequest() *AlibabaIdleIsvRefundQueryRequest{
    return &AlibabaIdleIsvRefundQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvRefundQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.refund.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAppraiseIsvOrderQuery Setter
// 系统自动生成
func (r *AlibabaIdleIsvRefundQueryRequest) SetParamAppraiseIsvOrderQuery(_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
    r._paramAppraiseIsvOrderQuery = _paramAppraiseIsvOrderQuery
    r.Set("param_appraise_isv_order_query", _paramAppraiseIsvOrderQuery)
    return nil
}

// ParamAppraiseIsvOrderQuery Getter
func (r AlibabaIdleIsvRefundQueryRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
    return r._paramAppraiseIsvOrderQuery
}
