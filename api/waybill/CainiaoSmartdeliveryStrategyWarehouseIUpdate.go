package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoSmartdeliveryStrategyWarehouseIUpdate 智能发货引擎策略仓设置
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
func CainiaoSmartdeliveryStrategyWarehouseIUpdate(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest, resp *waybill.CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
