package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionordercreate 商旅机票分销-创建订单
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
func Alitripbtripflightdistributionordercreate(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionordercreateAPIRequest, session string) (*btrip.AlitripbtripflightdistributionordercreateAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
