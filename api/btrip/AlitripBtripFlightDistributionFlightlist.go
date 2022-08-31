package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionFlightlist 商旅机票航班列表接口
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
func AlitripBtripFlightDistributionFlightlist(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionFlightlistAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionFlightlistAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionFlightlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
