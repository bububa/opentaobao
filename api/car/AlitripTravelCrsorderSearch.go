package car

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripTravelCrsorderSearch CRS接送机订单列表搜索
// alitrip.travel.crsorder.search
//
// 提供给CRS商家搜索订单列表，仅返回订单号列表
func AlitripTravelCrsorderSearch(ctx context.Context, clt *core.SDKClient, req *car.AlitripTravelCrsorderSearchAPIRequest, resp *car.AlitripTravelCrsorderSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
