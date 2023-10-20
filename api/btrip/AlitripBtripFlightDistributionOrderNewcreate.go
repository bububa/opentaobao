package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionordernewcreate 商旅机票分销-创建订单V2
// alitrip.btrip.flight.distribution.order.newcreate
//
// 商旅机票分销-创建订单V2
func Alitripbtripflightdistributionordernewcreate(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionordernewcreateAPIRequest, session string) (*btrip.AlitripbtripflightdistributionordernewcreateAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionordernewcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
