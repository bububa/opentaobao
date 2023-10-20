package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenSupplychainFlightTrade 【商旅】机票交易流水查询接口
// alitrip.btrip.open.supplychain.flight.trade
//
// 【商旅】杭州市政府机票交易流水接口查询
func AlitripBtripOpenSupplychainFlightTrade(clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainFlightTradeAPIRequest, resp *btrip.AlitripBtripOpenSupplychainFlightTradeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
