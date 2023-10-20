package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscItemActivityGet 查询无条件单品优惠活动
// taobao.promotionmisc.item.activity.get
//
// 查询无条件单品优惠活动
func TaobaoPromotionmiscItemActivityGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityGetAPIRequest, resp *promotion.TaobaoPromotionmiscItemActivityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
