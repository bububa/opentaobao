package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivityGet 查询营销活动
// taobao.ump.activity.get
//
// 查询营销活动
func TaobaoUmpActivityGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpActivityGetAPIRequest, resp *promotion.TaobaoUmpActivityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
