package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewdetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
func AlitripBtripFlightDistributionChangeNewdetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewdetailAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeNewdetailAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeNewdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
