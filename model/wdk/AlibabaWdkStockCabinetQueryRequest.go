package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库位库存查询 API请求
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息
*/
type AlibabaWdkStockCabinetQueryRequest struct {
    model.Params
    // 系统自动生成
    query   *WmsInventoryTopQuery
}

// 初始化AlibabaWdkStockCabinetQueryRequest对象
func NewAlibabaWdkStockCabinetQueryRequest() *AlibabaWdkStockCabinetQueryRequest{
    return &AlibabaWdkStockCabinetQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockCabinetQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.cabinet.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockCabinetQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 系统自动生成
func (r *AlibabaWdkStockCabinetQueryRequest) SetQuery(query *WmsInventoryTopQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaWdkStockCabinetQueryRequest) GetQuery() *WmsInventoryTopQuery {
    return r.query
}
