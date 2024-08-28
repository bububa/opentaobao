package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscActivityRangeAdd 增加活动参与的商品
// taobao.promotionmisc.activity.range.add
//
// 增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
func TaobaoPromotionmiscActivityRangeAdd(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeAddAPIRequest, resp *promotion.TaobaoPromotionmiscActivityRangeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
