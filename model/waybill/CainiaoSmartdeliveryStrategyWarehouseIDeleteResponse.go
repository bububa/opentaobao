package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除智能发货引擎仓策略 APIResponse
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略
*/
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse `json:"cainiao_smartdelivery_strategy_warehouse_i_delete_response,omitempty"` 
    CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse
}

/* model for simplify = false
type CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse struct {

    // data
    
    IsDeleteSuccess   bool `json:"is_delete_success,omitempty"`
    

}
*/

type CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse struct {

    // data
    IsDeleteSuccess   bool `json:"is_delete_success,omitempty"`

}
