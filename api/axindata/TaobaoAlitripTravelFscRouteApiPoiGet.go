package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiPoiGet 获取景点（POI）信息
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
func TaobaoAlitripTravelFscRouteApiPoiGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
