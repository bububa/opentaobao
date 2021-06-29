package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除智能发货引擎仓策略 API请求
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略
*/
type CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest struct {
    model.Params
    // 仓id
    warehouseId   int64
}

// 初始化CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest对象
func NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest() *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest{
    return &CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetApiMethodName() string {
    return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseId Setter
// 仓id
func (r *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetWarehouseId() int64 {
    return r.warehouseId
}
