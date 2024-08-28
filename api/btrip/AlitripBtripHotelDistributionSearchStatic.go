package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchStatic 商旅酒店api分销-酒店静态信息接口
// alitrip.btrip.hotel.distribution.search.static
//
// 商旅酒店api分销-酒店静态信息接口
func AlitripBtripHotelDistributionSearchStatic(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchStaticAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchStaticAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
