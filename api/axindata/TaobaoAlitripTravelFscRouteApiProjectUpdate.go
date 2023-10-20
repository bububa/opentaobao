package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectUpdate 更新团期
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
func TaobaoAlitripTravelFscRouteApiProjectUpdate(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProjectUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
