package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityDelete 删除通用单品优惠活动
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
func TaobaoPromotionmiscCommonItemActivityDelete(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
