package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpDetailUpdate 修改活动详情
// taobao.ump.detail.update
//
// 更新活动详情
func TaobaoUmpDetailUpdate(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpDetailUpdateAPIRequest, resp *promotion.TaobaoUmpDetailUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
