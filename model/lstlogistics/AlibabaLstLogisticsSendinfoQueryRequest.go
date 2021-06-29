package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-查询主订单包含的物流单 API请求
alibaba.lst.logistics.sendinfo.query

查询主订单包含的物流单
*/
type AlibabaLstLogisticsSendinfoQueryRequest struct {
    model.Params
    // 入参
    query   *LstLogisticsInfoQuery
}

// 初始化AlibabaLstLogisticsSendinfoQueryRequest对象
func NewAlibabaLstLogisticsSendinfoQueryRequest() *AlibabaLstLogisticsSendinfoQueryRequest{
    return &AlibabaLstLogisticsSendinfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsSendinfoQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.sendinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsSendinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *AlibabaLstLogisticsSendinfoQueryRequest) SetQuery(query *LstLogisticsInfoQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaLstLogisticsSendinfoQueryRequest) GetQuery() *LstLogisticsInfoQuery {
    return r.query
}
