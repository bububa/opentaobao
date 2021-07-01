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
type AlibabaWdkUmsInventoryAdjustGetAPIRequest struct {
    model.Params
    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    _warehouseCode   string
}

// 初始化AlibabaWdkUmsInventoryAdjustGetAPIRequest对象
func NewAlibabaWdkUmsInventoryAdjustGetRequest() *AlibabaWdkUmsInventoryAdjustGetAPIRequest{
    return &AlibabaWdkUmsInventoryAdjustGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inventory.adjust.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsInventoryAdjustGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
