package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductOnline 产品上线
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
func TaobaoAlitripTravelFscRouteApiProductOnline(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProductOnlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
