package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectClose 关闭团期
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
func TaobaoAlitripTravelFscRouteApiProjectClose(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
