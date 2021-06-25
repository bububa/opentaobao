package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库存查询 APIRequest
alibaba.wdk.stock.real.query

查询仓内实时库存信息
*/
type AlibabaWdkStockRealQueryRequest struct {
    model.Params

    // 系统自动生成
    query   *WmsInventoryTopQuery 

}

func NewAlibabaWdkStockRealQueryRequest() *AlibabaWdkStockRealQueryRequest{
    return &AlibabaWdkStockRealQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkStockRealQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.real.query"
}

func (r AlibabaWdkStockRealQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkStockRealQueryRequest) SetQuery(query *WmsInventoryTopQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaWdkStockRealQueryRequest) GetQuery() *WmsInventoryTopQuery {
    return r.query
}

