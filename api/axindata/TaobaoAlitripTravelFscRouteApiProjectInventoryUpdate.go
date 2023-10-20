package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectInventoryUpdate 更新团期库存
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
func TaobaoAlitripTravelFscRouteApiProjectInventoryUpdate(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
