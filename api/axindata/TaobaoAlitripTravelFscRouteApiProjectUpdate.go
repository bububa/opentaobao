package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectUpdate 更新团期
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
func TaobaoAlitripTravelFscRouteApiProjectUpdate(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
