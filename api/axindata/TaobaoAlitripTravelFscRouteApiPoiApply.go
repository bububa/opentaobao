package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiPoiApply 线路供应商提交新增POI申请
// taobao.alitrip.travel.fsc.route.api.poi.apply
//
// 线路供应商提交新增POI申请
func TaobaoAlitripTravelFscRouteApiPoiApply(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiPoiApplyAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiPoiApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
