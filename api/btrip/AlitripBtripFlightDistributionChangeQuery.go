package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeQuery 改签航班询价
// alitrip.btrip.flight.distribution.change.query
//
// 商旅机票改签航班询价
func AlitripBtripFlightDistributionChangeQuery(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeQueryAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeQueryAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
