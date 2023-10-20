package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductOffline 产品下线
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
func TaobaoAlitripTravelFscRouteApiProductOffline(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
