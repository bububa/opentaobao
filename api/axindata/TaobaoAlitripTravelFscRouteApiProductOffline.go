package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductOffline 产品下线
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
func TaobaoAlitripTravelFscRouteApiProductOffline(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
