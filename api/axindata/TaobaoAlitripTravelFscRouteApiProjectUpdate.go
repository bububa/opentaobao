package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectUpdate 更新团期
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
func TaobaoAlitripTravelFscRouteApiProjectUpdate(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
