package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectAdd 新增团期
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
func TaobaoAlitripTravelFscRouteApiProjectAdd(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
