package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenSupplychainHotelTrade 【商旅】酒店交易查询流水接口
// alitrip.btrip.open.supplychain.hotel.trade
//
// 【商旅】酒店交易查询流水接口——杭州市政府
func AlitripBtripOpenSupplychainHotelTrade(clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainHotelTradeAPIRequest, resp *btrip.AlitripBtripOpenSupplychainHotelTradeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
