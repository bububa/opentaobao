package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智能发货引擎策略仓设置 APIResponse
cainiao.smartdelivery.strategy.warehouse.i.update

智能发货引擎发货策略设置仓维度
*/
type CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse struct {
    model.CommonResponse
    CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse
}

type CainiaoSmartdeliveryStrategyWarehouseIUpdateResponse struct {
    XMLName xml.Name `xml:"cainiao_smartdelivery_strategy_warehouse_i_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 仓信息
    
    WarehouseInfo   *WarehouseDto `json:"warehouse_info,omitempty" xml:"warehouse_info,omitempty"`

    
}
