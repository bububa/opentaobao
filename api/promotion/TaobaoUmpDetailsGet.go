package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpDetailsGet 查询活动详情列表
// taobao.ump.details.get
//
// 分页查询优惠详情列表
func TaobaoUmpDetailsGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpDetailsGetAPIRequest, resp *promotion.TaobaoUmpDetailsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
