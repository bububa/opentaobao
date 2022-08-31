package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeCancel 商旅机票改签取消
// alitrip.btrip.flight.distribution.change.cancel
//
// 商旅机票改签取消
func AlitripBtripFlightDistributionChangeCancel(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeCancelAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeCancelAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
