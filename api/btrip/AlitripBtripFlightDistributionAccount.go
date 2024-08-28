package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionAccount 机票分销企业或者tmc企业预存or月结账户查询接口
// alitrip.btrip.flight.distribution.account
//
// 机票分销企业或者tmc企业预存or月结账户查询
func AlitripBtripFlightDistributionAccount(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionAccountAPIRequest, resp *btrip.AlitripBtripFlightDistributionAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
