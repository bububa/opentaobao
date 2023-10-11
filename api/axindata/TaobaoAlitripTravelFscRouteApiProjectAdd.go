package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProjectAdd 新增团期
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
func TaobaoAlitripTravelFscRouteApiProjectAdd(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProjectAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
