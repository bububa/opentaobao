package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductOnline 产品上线
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
func TaobaoAlitripTravelFscRouteApiProductOnline(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
