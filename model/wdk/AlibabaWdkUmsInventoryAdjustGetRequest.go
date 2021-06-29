package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库调单-回流单 API请求
alibaba.wdk.ums.inventory.adjust.get

库调单-回流单
*/
type AlibabaWdkUmsInventoryAdjustGetRequest struct {
    model.Params
    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    warehouseCode   string
}

// 初始化AlibabaWdkUmsInventoryAdjustGetRequest对象
func NewAlibabaWdkUmsInventoryAdjustGetRequest() *AlibabaWdkUmsInventoryAdjustGetRequest{
    return &AlibabaWdkUmsInventoryAdjustGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryAdjustGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inventory.adjust.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryAdjustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsInventoryAdjustGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsInventoryAdjustGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
