package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelProductBaseModify 供应商编辑产品API
// taobao.alitrip.travel.product.base.modify
//
// 飞猪供销平台供应商可通过该API编辑产品
func TaobaoAlitripTravelProductBaseModify(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelProductBaseModifyAPIRequest, resp *travel.TaobaoAlitripTravelProductBaseModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
