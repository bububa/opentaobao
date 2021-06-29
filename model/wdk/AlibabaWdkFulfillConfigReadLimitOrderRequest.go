package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据仓code查询仓限单配置 API请求
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置
*/
type AlibabaWdkFulfillConfigReadLimitOrderRequest struct {
    model.Params
    // 仓code集合
    warehouseCodeList   []string
}

// 初始化AlibabaWdkFulfillConfigReadLimitOrderRequest对象
func NewAlibabaWdkFulfillConfigReadLimitOrderRequest() *AlibabaWdkFulfillConfigReadLimitOrderRequest{
    return &AlibabaWdkFulfillConfigReadLimitOrderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillConfigReadLimitOrderRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.config.read.limit.order"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillConfigReadLimitOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCodeList Setter
// 仓code集合
func (r *AlibabaWdkFulfillConfigReadLimitOrderRequest) SetWarehouseCodeList(warehouseCodeList []string) error {
    r.warehouseCodeList = warehouseCodeList
    r.Set("warehouse_code_list", warehouseCodeList)
    return nil
}

// WarehouseCodeList Getter
func (r AlibabaWdkFulfillConfigReadLimitOrderRequest) GetWarehouseCodeList() []string {
    return r.warehouseCodeList
}
