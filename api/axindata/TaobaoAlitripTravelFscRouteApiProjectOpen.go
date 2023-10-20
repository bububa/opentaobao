package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectOpen 打开团期
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
func TaobaoAlitripTravelFscRouteApiProjectOpen(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectOpenAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProjectOpenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
