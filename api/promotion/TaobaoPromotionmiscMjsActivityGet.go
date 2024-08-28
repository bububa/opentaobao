package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscMjsActivityGet 查询满就送活动
// taobao.promotionmisc.mjs.activity.get
//
// 查询满就送活动
func TaobaoPromotionmiscMjsActivityGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscMjsActivityGetAPIRequest, resp *promotion.TaobaoPromotionmiscMjsActivityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
