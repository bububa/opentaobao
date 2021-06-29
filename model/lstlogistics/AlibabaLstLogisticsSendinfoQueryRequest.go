package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-查询主订单包含的物流单 APIRequest
alibaba.lst.logistics.sendinfo.query

查询主订单包含的物流单
*/
type AlibabaLstLogisticsSendinfoQueryRequest struct {
    model.Params

    // 入参
    query   *LstLogisticsInfoQuery 

}

func NewAlibabaLstLogisticsSendinfoQueryRequest() *AlibabaLstLogisticsSendinfoQueryRequest{
    return &AlibabaLstLogisticsSendinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstLogisticsSendinfoQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.sendinfo.query"
}

func (r AlibabaLstLogisticsSendinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstLogisticsSendinfoQueryRequest) SetQuery(query *LstLogisticsInfoQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstLogisticsSendinfoQueryRequest) GetQuery() *LstLogisticsInfoQuery {
    return r.query
}

