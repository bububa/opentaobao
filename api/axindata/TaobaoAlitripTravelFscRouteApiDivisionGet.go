package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiDivisionGet 获取国家城市信息
// taobao.alitrip.travel.fsc.route.api.division.get
//
// 获取国家城市信息
func TaobaoAlitripTravelFscRouteApiDivisionGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
