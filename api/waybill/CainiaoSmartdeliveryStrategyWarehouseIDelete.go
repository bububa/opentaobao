package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaosmartdeliverystrategywarehouseidelete 删除智能发货引擎仓策略
// cainiao.smartdelivery.strategy.warehouse.i.delete
//
// 删除智能发货引擎仓策略
func Cainiaosmartdeliverystrategywarehouseidelete(clt *core.SDKClient, req *waybill.CainiaosmartdeliverystrategywarehouseideleteAPIRequest, session string) (*waybill.CainiaosmartdeliverystrategywarehouseideleteAPIResponse, error) {
	var resp waybill.CainiaosmartdeliverystrategywarehouseideleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
