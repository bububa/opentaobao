package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionFlightlist 商旅机票航班列表接口
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
func AlitripBtripFlightDistributionFlightlist(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionFlightlistAPIRequest, resp *btrip.AlitripBtripFlightDistributionFlightlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
