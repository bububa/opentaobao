package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoSmartdeliveryStrategyWarehouseIDelete 删除智能发货引擎仓策略
// cainiao.smartdelivery.strategy.warehouse.i.delete
//
// 删除智能发货引擎仓策略
func CainiaoSmartdeliveryStrategyWarehouseIDelete(clt *core.SDKClient, req *waybill.CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest, session string) (*waybill.CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse, error) {
	var resp waybill.CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
