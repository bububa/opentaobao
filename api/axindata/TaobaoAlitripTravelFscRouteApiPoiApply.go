package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiPoiApply 线路供应商提交新增POI申请
// taobao.alitrip.travel.fsc.route.api.poi.apply
//
// 线路供应商提交新增POI申请
func TaobaoAlitripTravelFscRouteApiPoiApply(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
