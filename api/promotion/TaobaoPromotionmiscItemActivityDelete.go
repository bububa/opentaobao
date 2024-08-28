package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscItemActivityDelete 删除无条件单品优惠活动
// taobao.promotionmisc.item.activity.delete
//
// 删除无条件单品优惠活动
func TaobaoPromotionmiscItemActivityDelete(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityDeleteAPIRequest, resp *promotion.TaobaoPromotionmiscItemActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
