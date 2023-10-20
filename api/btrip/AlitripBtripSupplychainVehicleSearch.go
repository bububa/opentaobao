package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainVehicleSearch 【商旅】用车订单搜索
// alitrip.btrip.supplychain.vehicle.search
//
// 【商旅】用车订单搜索
func AlitripBtripSupplychainVehicleSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainVehicleSearchAPIRequest, resp *btrip.AlitripBtripSupplychainVehicleSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
