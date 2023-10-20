package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewdetail 商旅机票退票详情接口V2
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
func AlitripBtripFlightDistributionRefundNewdetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewdetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundNewdetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
