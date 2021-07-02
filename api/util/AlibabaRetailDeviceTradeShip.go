package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaRetailDeviceTradeShip 贩卖机掉货成功通知
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
func AlibabaRetailDeviceTradeShip(clt *core.SDKClient, req *util.AlibabaRetailDeviceTradeShipAPIRequest, session string) (*util.AlibabaRetailDeviceTradeShipAPIResponse, error) {
	var resp util.AlibabaRetailDeviceTradeShipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
