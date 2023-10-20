package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionaccount 机票分销企业或者tmc企业预存or月结账户查询接口
// alitrip.btrip.flight.distribution.account
//
// 机票分销企业或者tmc企业预存or月结账户查询
func Alitripbtripflightdistributionaccount(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionaccountAPIRequest, session string) (*btrip.AlitripbtripflightdistributionaccountAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
