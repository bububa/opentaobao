package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据仓code查询仓限单配置 APIRequest
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置
*/
type AlibabaWdkFulfillConfigReadLimitOrderRequest struct {
    model.Params

    // 仓code集合
    warehouseCodeList   []String 

}

func NewAlibabaWdkFulfillConfigReadLimitOrderRequest() *AlibabaWdkFulfillConfigReadLimitOrderRequest{
    return &AlibabaWdkFulfillConfigReadLimitOrderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillConfigReadLimitOrderRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.config.read.limit.order"
}

func (r AlibabaWdkFulfillConfigReadLimitOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillConfigReadLimitOrderRequest) SetWarehouseCodeList(warehouseCodeList []String) error {
    r.warehouseCodeList = warehouseCodeList
    r.Set("warehouse_code_list", warehouseCodeList)
    return nil
}

func (r AlibabaWdkFulfillConfigReadLimitOrderRequest) GetWarehouseCodeList() []String {
    return r.warehouseCodeList
}

