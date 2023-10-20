package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightordersearch 获取机票订单列表
// alitrip.btrip.flight.order.search
//
// 第三方OA厂商获取机票订单列表
func Alitripbtripflightordersearch(clt *core.SDKClient, req *btrip.AlitripbtripflightordersearchAPIRequest, session string) (*btrip.AlitripbtripflightordersearchAPIResponse, error) {
	var resp btrip.AlitripbtripflightordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
