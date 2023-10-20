package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenSupplychainVehicleTrade 商旅用车交易流水接口
// alitrip.btrip.open.supplychain.vehicle.trade
//
// 商旅用车交易流水接口
func AlitripBtripOpenSupplychainVehicleTrade(clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainVehicleTradeAPIRequest, resp *btrip.AlitripBtripOpenSupplychainVehicleTradeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
