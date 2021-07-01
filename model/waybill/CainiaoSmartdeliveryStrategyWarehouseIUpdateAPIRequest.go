package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest
智能发货引擎策略仓设置 API请求
cainiao.smartdelivery.strategy.warehouse.i.update

智能发货引擎发货策略设置仓维度 */
type CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest struct {
	model.Params
	// 智能发货设置请求参数
	_deliveryStrategySetRequest *DeliveryStrategySetRequest
}

// New
