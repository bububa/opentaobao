package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoSmartdeliveryStrategyWarehouseIUpdate 智能发货引擎策略仓设置
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
func CainiaoSmartdeliveryStrategyWarehouseIUpdate(clt *core.SDKClient, req *waybill.CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest, session string) (*waybill.CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse, error) {
	var resp waybill.CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
