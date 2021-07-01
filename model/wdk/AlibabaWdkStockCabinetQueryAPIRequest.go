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
type AlibabaWdkStockCabinetQueryAPIRequest struct {
    model.Params
    // 系统自动生成
    _query   *WmsInventoryTopQuery
}

// 初始化AlibabaWdkStockCabinetQueryAPIRequest对象
func NewAlibabaWdkStockCabinetQueryRequest() *AlibabaWdkStockCabinetQueryAPIRequest{
    return &AlibabaWdkStockCabinetQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockCabinetQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.cabinet.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockCabinetQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 系统自动生成
func (r *AlibabaWdkStockCabinetQueryAPIRequest) SetQuery(_query *WmsInventoryTopQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaWdkStockCabinetQueryAPIRequest) GetQuery() *WmsInventoryTopQuery {
    return r._query
}
