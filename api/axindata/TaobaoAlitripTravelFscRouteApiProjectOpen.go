package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectOpen 打开团期
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
func TaobaoAlitripTravelFscRouteApiProjectOpen(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
