package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest
删除智能发货引擎仓策略 API请求
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略 */
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest struct {
	model.Params
	// 仓id
	_warehouseId int64
}

// New
