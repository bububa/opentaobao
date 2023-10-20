package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaosmartdeliverystrategywarehouseiupdateAPIResponse 智能发货引擎策略仓设置 API返回值
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
type CainiaosmartdeliverystrategywarehouseiupdateAPIResponse struct {
	model.CommonResponse
	CainiaosmartdeliverystrategywarehouseiupdateAPIResponseModel
}

// CainiaosmartdeliverystrategywarehouseiupdateAPIResponseModel is 智能发货引擎策略仓设置 成功返回结果
type CainiaosmartdeliverystrategywarehouseiupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_smartdelivery_strategy_warehouse_i_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓信息
	WarehouseInfo *WarehouseDto `json:"warehouse_info,omitempty" xml:"warehouse_info,omitempty"`
}
