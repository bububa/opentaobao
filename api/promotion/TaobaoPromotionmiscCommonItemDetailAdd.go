package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemDetailAdd 创建通用单品优惠详情
// taobao.promotionmisc.common.item.detail.add
//
// 创建通用单品优惠详情。
// 1、使用此接口在指定的优惠活动下创建参与的商品的优惠信息，如还未创建活动，需要先使用接口taobao.promotionmisc.common.item.activity.add创建优惠活动；
// 2、同一卖家同一活动下的优惠详情数量限制为150个，超过限制需先调用taobao.promotionmisc.common.item.detail.delete接口删除无用的详情后才可再创建新的优惠详情；
// 3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能创建
func TaobaoPromotionmiscCommonItemDetailAdd(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemDetailAddAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemDetailAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
