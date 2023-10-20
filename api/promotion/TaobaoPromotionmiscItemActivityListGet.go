package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscItemActivityListGet 查询无条件单品优惠活动列表
// taobao.promotionmisc.item.activity.list.get
//
// 查询无条件单品优惠活动列表
func TaobaoPromotionmiscItemActivityListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityListGetAPIRequest, resp *promotion.TaobaoPromotionmiscItemActivityListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
