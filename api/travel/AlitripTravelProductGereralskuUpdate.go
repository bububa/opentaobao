package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripTravelProductGereralskuUpdate (供销)船票通用类目sku新增&编辑API
// alitrip.travel.product.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
func AlitripTravelProductGereralskuUpdate(ctx context.Context, clt *core.SDKClient, req *travel.AlitripTravelProductGereralskuUpdateAPIRequest, resp *travel.AlitripTravelProductGereralskuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
