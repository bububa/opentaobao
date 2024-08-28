package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductAdd 新增线路产品基本信息
// taobao.alitrip.travel.fsc.route.api.product.add
//
// 新增线路产品基本信息
func TaobaoAlitripTravelFscRouteApiProductAdd(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductAddAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
