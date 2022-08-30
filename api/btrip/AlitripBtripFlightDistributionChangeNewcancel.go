package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewcancel 商旅机票改签取消
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
func AlitripBtripFlightDistributionChangeNewcancel(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewcancelAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeNewcancelAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeNewcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
