package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-查询运单物流追踪信息 API请求
alibaba.lst.logistics.trace.query

查询LP单物流追踪信息
*/
type AlibabaLstLogisticsTraceQueryRequest struct {
    model.Params
    // 入参
    _query   *LstLogisticsTraceQuery
}

// 初始化AlibabaLstLogisticsTraceQueryRequest对象
func NewAlibabaLstLogisticsTraceQueryRequest() *AlibabaLstLogisticsTraceQueryRequest{
    return &AlibabaLstLogisticsTraceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsTraceQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.trace.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsTraceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *AlibabaLstLogisticsTraceQueryRequest) SetQuery(_query *LstLogisticsTraceQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaLstLogisticsTraceQueryRequest) GetQuery() *LstLogisticsTraceQuery {
    return r._query
}
