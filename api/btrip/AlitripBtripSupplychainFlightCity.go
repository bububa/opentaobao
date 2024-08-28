package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainFlightCity 机场数据查询
// alitrip.btrip.supplychain.flight.city
//
// 机场数据查询
func AlitripBtripSupplychainFlightCity(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightCityAPIRequest, resp *btrip.AlitripBtripSupplychainFlightCityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
