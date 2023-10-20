package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewquery 改签航班询价V2
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
func AlitripBtripFlightDistributionChangeNewquery(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewqueryAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeNewqueryAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeNewqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
