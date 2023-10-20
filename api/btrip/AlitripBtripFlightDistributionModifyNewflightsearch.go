package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionModifyNewflightsearch 改签航班列表V2
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
func AlitripBtripFlightDistributionModifyNewflightsearch(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
