package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionModifyFlightsearch 改签航班列表
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
func AlitripBtripFlightDistributionModifyFlightsearch(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIRequest, resp *btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
