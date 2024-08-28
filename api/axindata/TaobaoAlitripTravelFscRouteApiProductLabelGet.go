package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductLabelGet 获取线路主题
// taobao.alitrip.travel.fsc.route.api.product.label.get
//
// 获取线路主题
func TaobaoAlitripTravelFscRouteApiProductLabelGet(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
