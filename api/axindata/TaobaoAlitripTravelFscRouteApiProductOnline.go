package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductOnline 产品上线
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
func TaobaoAlitripTravelFscRouteApiProductOnline(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
