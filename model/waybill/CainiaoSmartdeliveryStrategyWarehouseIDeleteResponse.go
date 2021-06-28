package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除智能发货引擎仓策略 APIResponse
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略
*/
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_smartdelivery_strategy_warehouse_i_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    IsDeleteSuccess   bool `json:"is_delete_success,omitempty" xml:"