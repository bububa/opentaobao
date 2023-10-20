package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiDivisionGet 获取国家城市信息
// taobao.alitrip.travel.fsc.route.api.division.get
//
// 获取国家城市信息
func TaobaoAlitripTravelFscRouteApiDivisionGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiDivisionGetAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiDivisionGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
