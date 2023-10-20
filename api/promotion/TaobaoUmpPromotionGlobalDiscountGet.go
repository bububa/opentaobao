package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpPromotionGlobalDiscountGet 获取卖家最低折扣
// taobao.ump.promotion.global.discount.get
//
// 提供卖家最低折扣查询功能
func TaobaoUmpPromotionGlobalDiscountGet(clt *core.SDKClient, req *promotion.TaobaoUmpPromotionGlobalDiscountGetAPIRequest, resp *promotion.TaobaoUmpPromotionGlobalDiscountGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
