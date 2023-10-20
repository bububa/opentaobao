package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionAccount 机票分销企业或者tmc企业预存or月结账户查询接口
// alitrip.btrip.flight.distribution.account
//
// 机票分销企业或者tmc企业预存or月结账户查询
func AlitripBtripFlightDistributionAccount(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionAccountAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionAccountAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
