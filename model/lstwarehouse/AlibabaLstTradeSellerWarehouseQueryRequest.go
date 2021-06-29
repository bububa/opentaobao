package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-本云商家-仓库查询接口 APIRequest
alibaba.lst.trade.seller.warehouse.query

查询本地云仓商家的仓库
*/
type AlibabaLstTradeSellerWarehouseQueryRequest struct {
    model.Params

    // 入参
    warehouseQueryParam   *WarehouseQueryParam 

}

func NewAlibabaLstTradeSellerWarehouseQueryRequest() *AlibabaLstTradeSellerWarehouseQueryRequest{
    return &AlibabaLstTradeSellerWarehouseQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerWarehouseQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.warehouse.query"
}

func (r AlibabaLstTradeSellerWarehouseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerWarehouseQueryRequest) SetWarehouseQueryParam(warehouseQueryParam *WarehouseQueryParam) error {
    r.warehouseQueryParam = warehouseQueryParam
    r.Set("warehouse_query_param", warehouseQueryParam)
    return nil
}

func (r AlibabaLstTradeSellerWarehouseQueryRequest) GetWarehouseQueryParam() *WarehouseQueryParam {
    return r.warehouseQueryParam
}

