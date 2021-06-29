package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-查询运单物流追踪信息 APIRequest
alibaba.lst.logistics.trace.query

查询LP单物流追踪信息
*/
type AlibabaLstLogisticsTraceQueryRequest struct {
    model.Params

    // 入参
    query   *LstLogisticsTraceQuery 

}

func NewAlibabaLstLogisticsTraceQueryRequest() *AlibabaLstLogisticsTraceQueryRequest{
    return &AlibabaLstLogisticsTraceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstLogisticsTraceQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.trace.query"
}

func (r AlibabaLstLogisticsTraceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstLogisticsTraceQueryRequest) SetQuery(query *LstLogisticsTraceQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstLogisticsTraceQueryRequest) GetQuery() *LstLogisticsTraceQuery {
    return r.query
}

