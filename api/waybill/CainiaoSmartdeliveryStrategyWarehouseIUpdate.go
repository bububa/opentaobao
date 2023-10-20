package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaosmartdeliverystrategywarehouseiupdate 智能发货引擎策略仓设置
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
func Cainiaosmartdeliverystrategywarehouseiupdate(clt *core.SDKClient, req *waybill.CainiaosmartdeliverystrategywarehouseiupdateAPIRequest, session string) (*waybill.CainiaosmartdeliverystrategywarehouseiupdateAPIResponse, error) {
	var resp waybill.CainiaosmartdeliverystrategywarehouseiupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
