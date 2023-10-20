package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiBusinessAreaGet 获取业务区域
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
func TaobaoAlitripTravelFscRouteApiBusinessAreaGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
