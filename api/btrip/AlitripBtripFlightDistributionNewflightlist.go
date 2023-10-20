package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionNewflightlist 商旅机票航班列表接口，用于分销询价V2
// alitrip.btrip.flight.distribution.newflightlist
//
// 商旅机票航班列表接口，用于分销询价V2
func AlitripBtripFlightDistributionNewflightlist(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionNewflightlistAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionNewflightlistAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionNewflightlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
