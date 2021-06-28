package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
智能发货引擎策略仓设置 APIResponse
cainiao.smartdelivery.strategy.warehouse.i.update

智能发货引擎发货策略设置仓维度
*/
type CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse `json:"cainiao_smartdelivery_strategy_warehouse_i_update_response,omitempty"` 
    CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse
}

/* model for simplify = false
type CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse struct {

    // 仓信息
    
    WarehouseInfo  *struct {
        WarehouseDto  *WarehouseDto `json:"warehouse_dto,omitempty"`
    } `json:"warehouse_info,omitempty"`
    

}
*/

type CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse struct {

    // 仓信息
    WarehouseInfo   *WarehouseDto `json:"warehouse_info,omitempty"`

}
