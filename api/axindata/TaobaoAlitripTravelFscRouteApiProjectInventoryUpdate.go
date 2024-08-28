package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectInventoryUpdate 更新团期库存
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
func TaobaoAlitripTravelFscRouteApiProjectInventoryUpdate(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
