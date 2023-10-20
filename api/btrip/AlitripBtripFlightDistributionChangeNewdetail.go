package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewdetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
func AlitripBtripFlightDistributionChangeNewdetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewdetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewdetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
