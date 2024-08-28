package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripTravelGereralskuUpdate 发布SKU信息（如果properties重复 则更新）
// alitrip.travel.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
func AlitripTravelGereralskuUpdate(ctx context.Context, clt *core.SDKClient, req *travel.AlitripTravelGereralskuUpdateAPIRequest, resp *travel.AlitripTravelGereralskuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
