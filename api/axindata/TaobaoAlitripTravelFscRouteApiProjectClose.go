package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectClose 关闭团期
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
func TaobaoAlitripTravelFscRouteApiProjectClose(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectCloseAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProjectCloseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
