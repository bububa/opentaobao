package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeDetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.detail
//
// 商旅机票改签详情接口
func AlitripBtripFlightDistributionChangeDetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeDetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
