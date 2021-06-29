package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-本云商家-仓库查询接口 API请求
alibaba.lst.trade.seller.warehouse.query

查询本地云仓商家的仓库
*/
type AlibabaLstTradeSellerWarehouseQueryRequest struct {
    model.Params
    // 入参
    warehouseQueryParam   *WarehouseQueryParam
}

// 初始化AlibabaLstTradeSellerWarehouseQueryRequest对象
func NewAlibabaLstTradeSellerWarehouseQueryRequest() *AlibabaLstTradeSellerWarehouseQueryRequest{
    return &AlibabaLstTradeSellerWarehouseQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerWarehouseQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.warehouse.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerWarehouseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseQueryParam Setter
// 入参
func (r *AlibabaLstTradeSellerWarehouseQueryRequest) SetWarehouseQueryParam(warehouseQueryParam *WarehouseQueryParam) error {
    r.warehouseQueryParam = warehouseQueryParam
    r.Set("warehouse_query_param", warehouseQueryParam)
    return nil
}

// WarehouseQueryParam Getter
func (r AlibabaLstTradeSellerWarehouseQueryRequest) GetWarehouseQueryParam() *WarehouseQueryParam {
    return r.warehouseQueryParam
}
