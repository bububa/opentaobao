package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库位库存查询 APIRequest
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息
*/
type AlibabaWdkStockCabinetQueryRequest struct {
    model.Params

    // 系统自动生成
    query   *WmsInventoryTopQuery 

}

func NewAlibabaWdkStockCabinetQueryRequest() *AlibabaWdkStockCabinetQueryRequest{
    return &AlibabaWdkStockCabinetQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkStockCabinetQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.cabinet.query"
}

func (r AlibabaWdkStockCabinetQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkStockCabinetQueryRequest) SetQuery(query *WmsInventoryTopQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaWdkStockCabinetQueryRequest) GetQuery() *WmsInventoryTopQuery {
    return r.query
}

