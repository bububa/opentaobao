package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelProductSkuOverride （供销）产品级别日历价格库存修改，全量覆盖
// taobao.alitrip.travel.product.sku.override
//
// （供销）产品级别日历价格库存修改，全量覆盖
func TaobaoAlitripTravelProductSkuOverride(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelProductSkuOverrideAPIRequest, resp *travel.TaobaoAlitripTravelProductSkuOverrideAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
