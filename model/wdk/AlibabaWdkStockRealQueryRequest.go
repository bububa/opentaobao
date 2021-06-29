package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库存查询 API请求
alibaba.wdk.stock.real.query

查询仓内实时库存信息
*/
type AlibabaWdkStockRealQueryRequest struct {
    model.Params
    // 系统自动生成
    _query   *WmsInventoryTopQuery
}

// 初始化AlibabaWdkStockRealQueryRequest对象
func NewAlibabaWdkStockRealQueryRequest() *AlibabaWdkStockRealQueryRequest{
    return &AlibabaWdkStockRealQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockRealQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.real.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockRealQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 系统自动生成
func (r *AlibabaWdkStockRealQueryRequest) SetQuery(_query *WmsInventoryTopQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaWdkStockRealQueryRequest) GetQuery() *WmsInventoryTopQuery {
    return r._query
}
