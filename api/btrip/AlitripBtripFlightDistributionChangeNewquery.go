package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewquery 改签航班询价V2
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
func AlitripBtripFlightDistributionChangeNewquery(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewqueryAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewqueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
