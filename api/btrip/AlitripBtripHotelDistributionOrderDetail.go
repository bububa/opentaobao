package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionorderdetail 商旅酒店API分销查询订单详情
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
func Alitripbtriphoteldistributionorderdetail(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionorderdetailAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionorderdetailAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionorderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
