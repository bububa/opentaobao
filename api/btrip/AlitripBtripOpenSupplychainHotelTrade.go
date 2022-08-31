package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenSupplychainHotelTrade 【商旅】酒店交易查询流水接口
// alitrip.btrip.open.supplychain.hotel.trade
//
// 【商旅】酒店交易查询流水接口——杭州市政府
func AlitripBtripOpenSupplychainHotelTrade(clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainHotelTradeAPIRequest, session string) (*btrip.AlitripBtripOpenSupplychainHotelTradeAPIResponse, error) {
	var resp btrip.AlitripBtripOpenSupplychainHotelTradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
