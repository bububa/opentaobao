package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityGet 查询通用单品优惠活动
// taobao.promotionmisc.common.item.activity.get
//
// 查询通用单品优惠活动。
func TaobaoPromotionmiscCommonItemActivityGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityGetAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemActivityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
