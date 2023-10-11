package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiPoiGet 获取景点（POI）信息
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
func TaobaoAlitripTravelFscRouteApiPoiGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
