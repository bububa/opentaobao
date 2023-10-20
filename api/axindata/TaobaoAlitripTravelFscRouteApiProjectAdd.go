package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectAdd 新增团期
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
func TaobaoAlitripTravelFscRouteApiProjectAdd(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
