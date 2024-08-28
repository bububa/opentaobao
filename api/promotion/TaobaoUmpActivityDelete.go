package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivityDelete 删除营销活动
// taobao.ump.activity.delete
//
// 删除营销活动。对应的活动详情等将会被全部删除。
func TaobaoUmpActivityDelete(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpActivityDeleteAPIRequest, resp *promotion.TaobaoUmpActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
