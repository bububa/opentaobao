package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionModifyFlightsearch 改签航班列表
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
func AlitripBtripFlightDistributionModifyFlightsearch(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
