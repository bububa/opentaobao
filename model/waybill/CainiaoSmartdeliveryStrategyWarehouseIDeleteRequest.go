package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除智能发货引擎仓策略 APIRequest
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略
*/
type CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest struct {
    model.Params

    // 仓id
    warehouseId   int64 

}

func NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest() *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest{
    return &CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetApiMethodName() string {
    return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest) GetWarehouseId() int64 {
    return r.warehouseId
}

